package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBName     string `mapstructure:"db_dbname"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBSslMode  string `mapstructure:"db_sslmode"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	// search config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	// confirm file is used
	log.Printf("Using config: %s\n", viper.ConfigFileUsed())
	// to unnarshal values into target object
	err = viper.Unmarshal(&config)
	return
	// loading function isc completed
}
