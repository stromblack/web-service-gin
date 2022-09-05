package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost      string `mapstructure:"db_host"`
	DBPort      string `mapstructure:"db_port"`
	DBName      string `mapstructure:"db_dbname"`
	DBUser      string `mapstructure:"db_user"`
	DBPassword  string `mapstructure:"db_password"`
	DBSslMode   string `mapstructure:"db_sslmode"`
	Secret      string `mapstructure:"jwt_secret"`
	Issuer      string `mapstructure:"jwt_iss"`
	Audience    string `mapstructure:"jwt_aud"`
	TokenExpire int    `mapstructure:"jwt_exp"`
}

func LoadConfig() (config Config, err error) {
	_, found := os.LookupEnv("GIN_ENV")
	var env string = "dev"
	if found {
		env = os.Getenv("GIN_ENV")
	}
	filename := fmt.Sprintf("app-%s", env)
	viper.SetConfigName(filename)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	// search config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	// confirm file is used
	fmt.Printf("# Using config: %s\n", viper.ConfigFileUsed())
	// to unnarshal values into target object
	err = viper.Unmarshal(&config)
	return
	// loading function isc completed
}

// const (
// 	db_host     = "127.0.0.1"
// 	db_port     = "5432"
// 	db_user     = "postgres"
// 	db_password = "3018"
// 	db_dbname   = "synergy_dev"
// 	db_sslmode  = "disable"
// 	jwt_secret  = "brown dog jumps over lazy fox"
// 	jwt_iss     = "web-service-gin"
// 	jwt_aud     = "https://localhost:43000"
// 	jwt_exp     = 1
// )

// func LoadConfig() (config Config, err error) {
// 	return Config{
// 		DBHost:      db_host,
// 		DBPort:      db_port,
// 		DBUser:      db_user,
// 		DBPassword:  db_password,
// 		DBSslMode:   db_sslmode,
// 		Secret:      jwt_secret,
// 		Issuer:      jwt_iss,
// 		Audience:    jwt_aud,
// 		TokenExpire: jwt_exp,
// 	}, nil
// }
