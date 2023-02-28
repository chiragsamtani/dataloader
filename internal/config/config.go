package config

import (
	"github.com/spf13/viper"
)

type ImmutableConfig interface {
	GetLogLevel()
	GetSupplierConfig()
}

type RootConfig struct {
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	SupplierConfig string `mapstructure:"SUPPLIER_CONFIG"`
}

func (rc *RootConfig) GetLogLevel() string {
	return rc.LogLevel
}

func (rc *RootConfig) GetSupplierConfig() string {
	return rc.SupplierConfig
}

func GetConfigFromEnv() (*RootConfig, error) {
	// use local config by default
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetConfigName("app.local")
	var config RootConfig
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}
