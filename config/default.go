package config

import "time"

const (
	DefaultDatabaseMaxIdleConns    = 3
	DefaultDatabaseMaxOpenConns    = 5
	DefaultDatabaseConnMaxLifetime = 1 * time.Hour
	DefaultDatabasePingInterval    = 5 * time.Second
	DefaultDatabaseRetryAttempts   = 3

	DefaultWorkerRetryAttempts = 3
	DefaultWorkerTaskRetention = 1 * time.Hour
	DefaultWorkerConcurrency   = 25
	DefaultWorkerNamespace     = "mazharul-islam"
	DefaultStateReindexTTL     = 2 * 24 * time.Hour // 2 days

	DefaultHTTPTimeout           = 100 * time.Second
	DefaultTLSHandshakeTimeout   = 50 * time.Second
	DefaultTLSInsecureSkipVerify = true
	DefaultHTTPPort              = "4000"
	DefaultSwaggerEndpoint       = "127.0.0.1:" + DefaultHTTPPort

	DefaultRedisLockDuration  = 5 * time.Second
	DefaultRedisRetryAttempts = 3
)
