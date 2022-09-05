package database

import (
	"database/sql"
	"fmt"
	"synergy/web-service-gin/common"

	"synergy/web-service-gin/common/config"

	. "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

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

func PrintStatementInfo(stmt SelectStatement) {
	query, args := stmt.Sql()

	fmt.Println("Parameterized query: ")
	fmt.Println("==============================")
	fmt.Println(query)
	fmt.Println("Arguments: ")
	fmt.Println(args)

	debugSQL := stmt.DebugSql()

	fmt.Println("\n\nDebug sql: ")
	fmt.Println("==============================")
	fmt.Println(debugSQL)
}
