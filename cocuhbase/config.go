package cocuhbase

import "time"

type Config struct {
	// The application username to Connect to the couchbase cluster.
	Username string
	// The application password to Connect to the couchbase cluster.
	Password string
	// The connection string for the couchbase cluster.
	Host string
	// The bucket to use for the couchbase cluster.
	Bucket string
	// The timeout for connecting to the couchbase cluster.
	ConnectionTimeout time.Duration
	// The timeout for KV operations.
	KVTimeout time.Duration
}

var ConfigDefault = Config{
	Username:          "admin",
	Password:          "123456",
	Host:              "127.0.0.1:8091",
	Bucket:            "fiber_storage",
	ConnectionTimeout: 3 * time.Second,
	KVTimeout:         1 * time.Second,
}

func configDefault(config ...Config) Config {
	if len(config) == 0 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Username == "" {
		cfg.Username = ConfigDefault.Username
	}
	if cfg.Password == "" {
		cfg.Password = ConfigDefault.Password
	}
	if cfg.Host == "" {
		cfg.Host = ConfigDefault.Host
	}
	if cfg.Bucket == "" {
		cfg.Bucket = ConfigDefault.Bucket
	}

	if cfg.ConnectionTimeout == 0 {
		cfg.ConnectionTimeout = ConfigDefault.ConnectionTimeout
	}

	if cfg.KVTimeout == 0 {
		cfg.KVTimeout = ConfigDefault.KVTimeout
	}

	return cfg
}
