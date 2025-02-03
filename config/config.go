package config

import (
	"github.com/spf13/viper"
)

// Stores all configuration of the application
type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBSource          string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	SwiftCSVPath      string `mapstructure:"SWIFT_CSV_PATH"`
	MigrationURL      string `mapstructure:"MIGRATION_URL"`
	GinMode           string `mapstructure:"GIN_MODE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
