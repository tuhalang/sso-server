package config

import (
	"github.com/spf13/viper"
	"github.com/tuhalang/authen/internal/logger"
)

// Config stores all the application configuration
type Config struct {
	RestServer RestServerConfig `mapstructure:"rest-server-config"`
	Tenants    []TenantConfig   `mapstructure:"tenant-config"`
}

// RestServerConfig stores all app rest server config
type RestServerConfig struct {
	Port       int    `mapstructure:"port"`
	Host       string `mapstructure:"host"`
	SslEnabled bool   `mapstructure:"ssl-enabled"`
}

// TenantConfig stores the tenant info
type TenantConfig struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	Database       string `mapstructure:"database"`
	DBDriver       string `mapstructure:"db-driver"`
	MaxOpenConns   int    `mapstructure:"max-open-conns"`
	MaxIdleConns   int    `mapstructure:"max-idle-conns"`
	ConnLifeTime   int    `mapstructure:"conn-life-time"`
	ConnIdleTime   int    `mapstructure:"conn-idle-time"`
	TenantName     string `mapstructure:"tenant-name"`
	JwtKey         string `mapstructure:"jwt-key"`
	JwtExpiredTime int    `mapstructure:"jwt-expired-time"`
}

type KeyStore struct {
	JwtKey string
	JwtExp int
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
