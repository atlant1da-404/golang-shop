package config

import (
	"fmt"
	"golang-shop/storage"

	"os"
)

type (
	Config struct {
		HTTP HTTPConfig
		Db   storage.DBConfig
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
	defaultSSLMode  = "disable"
)

func setFromEnv(cfg *Config) {
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")
	if cfg.HTTP.Port == "" {
		cfg.HTTP.Port = defaultPORT
	}

	cfg.Db.Host = os.Getenv("HTTP_DBHOST")
	if cfg.Db.Host == "" {
		cfg.HTTP.Host = defaultDBHOST
	}

	if cfg.Db.Port == "" {
		cfg.Db.Port = defaultDBPORT
	}

	cfg.Db.Name = os.Getenv("HTTP_DBNAME")
	if cfg.Db.Name == "" {
		cfg.Db.Name = defaultDBNAME
	}

	cfg.Db.User = os.Getenv("HTTP_DBUSER")
	if cfg.Db.User == "" {
		cfg.Db.User = defaultDBUSER
	}

	cfg.Db.Password = os.Getenv("HTTP_DBPASSWORD")
	if cfg.Db.Password == "" {
		cfg.Db.User = defaultPASSWORD
	}

	cfg.Db.SSLMode = os.Getenv("HTTP_DBSSLMODE")
	if cfg.Db.SSLMode == "" {
		cfg.Db.SSLMode = defaultSSLMode
	}
}

func (cfg *Config) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Db.Host, cfg.Db.User, cfg.Db.Password, cfg.Db.Name,
		cfg.Db.SSLMode,
	)
}
