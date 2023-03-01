package config

import "github.com/spf13/viper"

type ImmutableConfig interface {
	GetLogLevel()
	GetSupplierConfig()
	GetSqliteDsn()
	GetRepositoryType()
	GetDatabaseUsername() string
	GetDatabasePassword() string
	GetDatabaseHost() string
	GetDatabaseName() string
}

type RootConfig struct {
	LogLevel         string `mapstructure:"LOG_LEVEL"`
	SupplierConfig   string `mapstructure:"SUPPLIER_CONFIG"`
	RepositoryType   string `mapstructure:"REPOSITORY_TYPE"`
	DatabaseUsername string `mapstructure:"MYSQL_DATABASE_USERNAME"`
	DatabasePassword string `mapstructure:"MYSQL_DATABASE_PASSWORD"`
	DatabaseHost     string `mapstructure:"MYSQL_DATABASE_HOST"`
	DatabaseName     string `mapstructure:"MYSQL_DATABASE"`
}

func (rc *RootConfig) GetLogLevel() string {
	return rc.LogLevel
}

func (rc *RootConfig) GetSupplierConfig() string {
	return rc.SupplierConfig
}

func (rc *RootConfig) GetRepositoryType() string {
	return rc.RepositoryType
}

func (rc *RootConfig) GetDatabaseUsername() string {
	return rc.DatabaseUsername
}

func (rc *RootConfig) GetDatabasePassword() string {
	return rc.DatabasePassword
}

func (rc *RootConfig) GetDatabaseHost() string {
	return rc.DatabaseHost
}

func (rc *RootConfig) GetDatabaseName() string {
	return rc.DatabaseName
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
