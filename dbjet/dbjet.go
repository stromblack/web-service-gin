package dbjet

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	// dot import so that jet go code would resemble as much as native SQL
	// dot import is not mandatory

	"synergy/web-service-gin/common"
	"synergy/web-service-gin/postgres/public/model"
	"synergy/web-service-gin/postgres/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

const (
	host     = "synergy-dev.cioo1kbd0ezx.ap-southeast-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres "
	password = "synergy-dev"
	dbName   = "postgres "
)

func GetData() interface{} {
	// Connect to database
	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", connectString)
	common.CheckErr(err)
	panicOnError(err)
	defer db.Close()
	// Write query
	stmt := SELECT(table.Customers.AllColumns, table.Contacts.AllColumns).FROM(
		table.Customers.
			LEFT_JOIN(table.Contacts, table.Customers.CustomerID.EQ(table.Contacts.CustomerID)),
	)
	var dest []struct {
		model.Customers
		Contact []model.Contacts
	}
	err = stmt.Query(db, &dest)
	panicOnError(err)
	printStatementInfo(stmt)
	return dest
}
func printStatementInfo(stmt SelectStatement) {
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
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
