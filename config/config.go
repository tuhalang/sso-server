package config

import (
	"github.com/spf13/viper"
	"github.com/tuhalang/authen/internal/logger"
)

// Config stores all the application configuration
type Config struct {
	RestServer RestServerConfig `mapstructure:"rest-server-config"`
	Databases  []DatabaseConfig `mapstructure:"database-config"`
}

// RestServerConfig stores all app rest server config
type RestServerConfig struct {
	Port       int  `mapstructure:"port"`
	SslEnabled bool `mapstructure:"ssl-enabled"`
}

// DatabaseConfig stores the database connection info
type DatabaseConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	DBDriver  string `mapstructure:"db-driver"`
	SpaceName string `mapstructure:"space-name"`
}

// LoadConfig loads all config from file
func LoadConfig(configFile string) *Config {
	log := logger.Get()

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Cannot read config file!")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal().Err(err).Msg("Cannot unmarshal config file")
	}

	return &config
}
