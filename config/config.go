package config

import "os"

type (
	Config struct {
		HTTP HTTPConfig
	}

	HTTPConfig struct {
		Host string
		Port string
	}
)

func Init() *Config {
	var cfg Config
	setFromEnv(&cfg)
	return &cfg
}

const (
	defaultPORT     = "8081"
	defaultDBNAME   = "shop_db"
	defaultDBUSER   = "shop_user"
	defaultPASSWORD = "root1"
	defaultDBHOST   = "localhost"
	defaultDBPORT   = "5432"
)

func setFromEnv(cfg *Config) {
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")
	if cfg.HTTP.Port == "" {
		cfg.HTTP.Port = defaultPORT
	}
}
