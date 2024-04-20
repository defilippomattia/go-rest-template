package config

import (
	"database/sql"
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog"
)

type ServerDeps struct {
	Db     *sql.DB
	Logger zerolog.Logger
}

type Config struct {
	Server struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
	} `toml:"server"`

	DB struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Database string `toml:"database"`
	} `toml:"db"`

	Logging struct {
		Level string `toml:"level"`
	} `toml:"logging"`
}

func (p Config) String() string {
	return fmt.Sprintf("\n\tHost=%s, \n\tPort=%s, \n\tDBHost=%s, \n\tDBPort=%s, \n\tDBUser=%s, \n\tDBPassword=*************, \n\tDBDatabase=%s, \n\tLogLevel=%s",
		p.Server.Host, p.Server.Port, p.DB.Host, p.DB.Port, p.DB.User, p.DB.Database, p.Logging.Level)
}

func LoadConfig(w io.Writer, filePath string, config interface{}) error {
	fmt.Fprintf(w, "Loading config file: %s\n", filePath)
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}
	fmt.Fprintf(w, "Config loaded: %v\n", config)

	return nil

}

func ValidateConfig(cfg Config) error {
	if cfg.Server.Host == "" {
		return fmt.Errorf("server host is required")
	}
	if cfg.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	if cfg.DB.Host == "" {
		return fmt.Errorf("db host is required")
	}
	if cfg.DB.Port == "" {
		return fmt.Errorf("db port is required")
	}
	if cfg.DB.User == "" {
		return fmt.Errorf("db user is required")
	}
	if cfg.DB.Password == "" {
		return fmt.Errorf("db password is required")
	}
	if cfg.DB.Database == "" {
		return fmt.Errorf("db database is required")
	}
	debugLevels := []string{"debug", "info", "warn", "error"}

	if !contains(debugLevels, cfg.Logging.Level) {
		return fmt.Errorf("logging level must be one of: %v", debugLevels)
	}

	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
