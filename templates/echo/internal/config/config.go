package config

import "time"

type (
	Configuration struct {
		Application   Application   `mapstructure:"app"`
		PostgreSQL    PostgreSQL    `mapstructure:"postgresql"`
		Authorization Authorization `mapstructure:"authorization"`
		Hash          Hash          `mapstructure:"hash"`
	}

	Application struct {
		Name        string        `mapstructure:"name"`
		Version     string        `mapstructure:"version"`
		Environment string        `mapstructure:"environment"`
		Host        string        `mapstructure:"host"`
		Port        string        `mapstructure:"port"`
		Timeout     time.Duration `mapstructure:"timeout"`
		LogOption   string        `mapstructure:"log_option"`
		LogLevel    string        `mapstructure:"log_level"`
	}

	PostgreSQL struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
	}

	Authorization struct {
		BearerSecret    string        `mapstructure:"bearer_secret"`
		BearerDuration  time.Duration `mapstructure:"bearer_duration"`
		RefreshSecret   string        `mapstructure:"refresh_secret"`
		RefreshDuration time.Duration `mapstructure:"refresh_duration"`
	}

	Hash struct {
		Salt string `mapstructure:"salt"`
	}
)
