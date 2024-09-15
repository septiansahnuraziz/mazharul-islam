package config

import (
	"github.com/mazharul-islam/commons"
	"github.com/mazharul-islam/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.AddConfigPath("./../../..")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warningf("%v", err)
	}

	// Iterate through Viper configuration settings and set them as environment variables
	for _, k := range viper.AllKeys() {
		key := utils.StringToUpper(strings.Replace(k, ".", "_", -1))
		_ = os.Setenv(key, viper.GetString(k))
	}
}

func AppName() string {
	return viper.GetString("app.name")
}

func AppSlugName() string {
	return viper.GetString("app.slug_name")
}

func AppVersion() string {
	return viper.GetString("app.version")
}

func AppBuild() string {
	return viper.GetString("build")
}

func HTTPPort() string {
	return viper.GetString("port")
}

func EnvironmentMode() string {
	return viper.GetString("mode")
}

func DatabaseTimeZone() string {
	return viper.GetString("db.timezone")
}

func DatabaseUsername() string {
	return viper.GetString("db.user")
}

func DatabasePassword() string {
	return viper.GetString("db.password")
}

func DatabaseHost() string {
	return viper.GetString("db.host")
}

func DatabasePort() string {
	return viper.GetString("db.port")
}

func DatabaseName() string {
	return viper.GetString("db.name")
}

func DatabaseSSL() string {
	return viper.GetString("db.ssl_mode")
}

func DatabaseDSN() string {
	return utils.WriteStringTemplate(`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		DatabaseHost(),
		DatabaseUsername(),
		DatabasePassword(),
		DatabaseName(),
		DatabasePort(),
		DatabaseSSL(),
		DatabaseTimeZone(),
	)
}

func DatabaseMaxIdleConns() int {
	value := viper.GetInt("db.max_idle_conns")
	return utils.ValueOrDefault[int](value, DefaultDatabaseMaxIdleConns)
}

func DatabaseMaxOpenConns() int {
	value := viper.GetInt("db.max_open_conns")
	return utils.ValueOrDefault[int](value, DefaultDatabaseMaxOpenConns)
}

func DatabaseConnMaxLifetime() time.Duration {
	value := viper.GetString("db.conn_max_lifetime")
	return utils.ParseDurationWithDefault(value, DefaultDatabaseConnMaxLifetime)

}

func DatabaseRetryAttempts() float64 {
	value := viper.GetFloat64("db.retry_attempts")
	return utils.ValueOrDefault[float64](value, DefaultDatabaseRetryAttempts)
}

func DatabasePingInterval() time.Duration {
	value := viper.GetString("db.ping_interval")
	return utils.ParseDurationWithDefault(value, DefaultDatabasePingInterval)
}

func GetLogLevel() string {
	value := viper.GetString("log_level")
	return utils.ValueOrDefault[string](value, string(commons.LogLevelTrace))
}

func SwaggerEndpoint() string {
	value := viper.GetString("swagger.endpoint")
	return utils.ValueOrDefault[string](value, DefaultSwaggerEndpoint)
}

func SwaggerUsername() string {
	return viper.GetString("swagger.username")
}

func SwaggerPassword() string {
	return viper.GetString("swagger.password")
}

func RedisCacheHost() string {
	return viper.GetString("redis.cache_host")
}

func EnableCaching() bool {
	return viper.GetBool("enable_caching")
}

func RedisDialTimeout() time.Duration {
	return utils.ParseDurationWithDefault(viper.GetString("redis.dial_timeout"), 5*time.Second)
}

func RedisWriteTimeout() time.Duration {
	return utils.ParseDurationWithDefault(viper.GetString("redis.write_timeout"), 2*time.Second)
}

func RedisReadTimeout() time.Duration {
	return utils.ParseDurationWithDefault(viper.GetString("redis.read_timeout"), 2*time.Second)
}

func RedisMaxIdleConn() int {
	return utils.ValueOrDefault[int](utils.StringToInt[int](viper.GetString("redis.max_idle_conn")), 20)
}

func RedisMaxActiveConn() int {
	return utils.ValueOrDefault[int](utils.StringToInt[int](viper.GetString("redis.max_active_conn")), 50)
}
