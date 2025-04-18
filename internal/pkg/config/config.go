package config

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
)

var AppConfig Config

type Config struct {
	Port               int    `mapstructure:"PORT"`
	Environment        string `mapstructure:"ENVIRONMENT"`
	Debug              bool   `mapstructure:"DEBUG"`
	PgUrl              string `mapstructure:"PG_URL"`
	PgPoolMaxOpenConns int    `mapstructure:"PG_POOL_MAX_OPEN_CONNS"`
	PgPoolMaxIdleConns int    `mapstructure:"PG_POOL_MAX_IDLE_CONNS"`
	JWTRealm           string `mapstructure:"JWT_REALM"`
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	JWTExpired         int    `mapstructure:"JWT_EXPIRED"`
	JWTIssuer          string `mapstructure:"JWT_ISSUER"`
}

// IsValid checks if the config is valid
func (c Config) IsValid() bool {
	if c.Port == 0 || c.Environment == "" || c.JWTSecret == "" || c.JWTExpired == 0 || c.JWTRealm == "" || c.JWTIssuer == "" || c.PgUrl == "" {
		return false
	}

	return true
}

// InitializeAppConfig initializes the app config
func InitializeAppConfig(configPath string) error {
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/")
		viper.AddConfigPath("/")
	}
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		return err
	}

	if !AppConfig.IsValid() {
		return errors.New("invalid config")
	}
	if AppConfig.PgPoolMaxOpenConns == 0 {
		AppConfig.PgPoolMaxOpenConns = 10
	}
	if AppConfig.PgPoolMaxIdleConns == 0 {
		AppConfig.PgPoolMaxIdleConns = 5
	}
	return nil
}
