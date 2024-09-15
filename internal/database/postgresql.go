package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jpillora/backoff"
	"github.com/mazharul-islam/config"
	"github.com/mazharul-islam/utils"
	"regexp"
	"time"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	// PostgreSQL represents gorm PostgreSQL
	PostgreSQL *gorm.DB

	// StopTickerCh signal for closing ticker channel
	StopTickerCh chan bool

	sqlRegexp = regexp.MustCompile(`(\$\d+)|\?`)
)

func InitializePostgresConnection() (*gorm.DB, error) {
	fmt.Println("CONNECTION ", config.DatabaseDSN())
	conn, err := openPostgresConnection(config.DatabaseDSN())
	if err != nil {
		log.WithField("databaseDSN", config.DatabaseDSN()).Fatal("failed to connect postgresql database: ", err)
	}

	PostgreSQL = conn
	StopTickerCh = make(chan bool)

	go checkConnection(time.NewTicker(config.DatabasePingInterval()))

	PostgreSQL.Logger = NewGormCustomLogger()
	switch config.GetLogLevel() {
	case "error":
		PostgreSQL.Logger = PostgreSQL.Logger.LogMode(gormLogger.Error)
	case "warn":
		PostgreSQL.Logger = PostgreSQL.Logger.LogMode(gormLogger.Warn)
	default:
		PostgreSQL.Logger = PostgreSQL.Logger.LogMode(gormLogger.Info)
	}

	log.Info("Connection to PostgreSQL Server success...")

	return conn, nil
}

func openPostgresConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	conn.SetMaxIdleConns(config.DatabaseMaxIdleConns())
	conn.SetMaxOpenConns(config.DatabaseMaxOpenConns())
	conn.SetConnMaxLifetime(config.DatabaseConnMaxLifetime())

	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, err
	}

	return db, nil
}

func reconnectPostgresConnection() {
	b := backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    100 * time.Millisecond,
		Max:    1 * time.Second,
	}

	postgresRetryAttempts := config.DatabaseRetryAttempts()

	for b.Attempt() < postgresRetryAttempts {
		conn, err := openPostgresConnection(config.DatabaseDSN())
		if err != nil {
			log.WithField("databaseDSN", config.DatabaseDSN()).Error("failed to connect postgresql database: ", err)
		}

		if conn != nil {
			PostgreSQL = conn
			break
		}
		time.Sleep(b.Duration())
	}

	if b.Attempt() >= postgresRetryAttempts {
		log.Fatal("maximum retry to connect database")
	}

	b.Reset()
}

func checkConnection(ticker *time.Ticker) {
	for {
		select {
		case <-StopTickerCh:
			ticker.Stop()
			return
		case <-ticker.C:
			if _, err := PostgreSQL.DB(); err != nil {
				reconnectPostgresConnection()
			}
		}
	}
}

// GormCustomLogger override gorm logger
type GormCustomLogger struct {
	gormLogger.Config
}

// NewGormCustomLogger :nodoc:
func NewGormCustomLogger() *GormCustomLogger {
	return &GormCustomLogger{
		Config: gormLogger.Config{
			LogLevel:      gormLogger.Info,
			SlowThreshold: time.Second,
			Colorful:      true,
		},
	}
}

// LogMode :nodoc:
func (g *GormCustomLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	g.LogLevel = level
	return g
}

// Info :nodoc:
func (g *GormCustomLogger) Info(ctx context.Context, message string, values ...interface{}) {
	if g.LogLevel >= gormLogger.Info {
		log.WithContext(ctx).WithFields(log.Fields{"data": values}).Error(message)
	}
}

// Warn :nodoc:
func (g *GormCustomLogger) Warn(ctx context.Context, message string, values ...interface{}) {
	if g.LogLevel >= gormLogger.Warn {
		log.WithContext(ctx).WithFields(log.Fields{"data": values}).Warn(message)
	}

}

// Error :nodoc:
func (g *GormCustomLogger) Error(ctx context.Context, message string, values ...interface{}) {
	if g.LogLevel >= gormLogger.Error {
		log.WithContext(ctx).WithFields(log.Fields{"data": values}).Error(message)
	}
}

// Trace :nodoc:
func (g *GormCustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if g.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	logger := log.WithContext(ctx).WithFields(log.Fields{
		"took":       elapsed,
		"calledFrom": utils.MyCaller(4),
	})

	sqlLog := sqlRegexp.ReplaceAllString(sql, "%v")
	if rows >= 0 {
		logger.WithField("rows", rows)
	} else {
		logger.WithField("rows", "-")
	}

	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound) && g.LogLevel >= gormLogger.Error:
		logger.WithField("sql", sqlLog).Error(err)
	case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.LogLevel >= gormLogger.Warn:
		slowLog := utils.WriteStringTemplate("SLOW SQL >= %v", g.SlowThreshold)
		logger.WithField("sql", sqlLog).Warn(slowLog)
	case g.LogLevel >= gormLogger.Info:
		logger.Info(sqlLog)

	}
}
