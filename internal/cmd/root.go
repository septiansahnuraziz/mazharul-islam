package cmd

import (
	"github.com/mazharul-islam/config"
	"github.com/mazharul-islam/internal/database"
	"github.com/mazharul-islam/internal/logger"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var redisOptions *database.RedisConnectionPoolOptions

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-miscellaneous-service",
	Short: "go miscellaneous service console",
	Long:  `This is go miscellaneous service console`,
}

func init() {
	config.LoadConfig()
	logger.SetupLogger()

	redisOptions = &database.RedisConnectionPoolOptions{
		DialTimeout:     config.RedisDialTimeout(),
		ReadTimeout:     config.RedisReadTimeout(),
		WriteTimeout:    config.RedisWriteTimeout(),
		IdleCount:       config.RedisMaxIdleConn(),
		PoolSize:        config.RedisMaxActiveConn(),
		IdleTimeout:     240 * time.Second,
		MaxConnLifetime: 1 * time.Minute,
	}

	log.Info("Environment: ", config.EnvironmentMode())
}

// Execute :nodoc:
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
