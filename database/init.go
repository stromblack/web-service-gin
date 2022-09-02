package database

import (
	"database/sql"
	"fmt"
	"synergy/web-service-gin/common"

	"synergy/web-service-gin/common/config"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "127.0.0.1"
// 	port     = "5432"
// 	user     = "postgres"
// 	password = "3018"
// 	dbname   = "synergy_dev"
// 	sslmode  = "disable" // disable, require, verify-ca, verify-full
// )

func InitDb() (*sql.DB, error) {
	config, _ := config.LoadConfig()
	// setup format string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s \n", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSslMode)
	// connect
	db, err := sql.Open("postgres", connStr)
	// check error
	common.CheckErr(err)
	// return db.
	fmt.Println("# Database connect.")
	return db, err
}
