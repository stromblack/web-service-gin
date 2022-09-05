//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var TUserInfo = newTUserInfoTable("public", "T_USER_INFO", "")

type tUserInfoTable struct {
	postgres.Table

	//Columns
	UID      postgres.ColumnInteger
	Username postgres.ColumnString
	Password postgres.ColumnString
	Email    postgres.ColumnString
	Created  postgres.ColumnDate

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TUserInfoTable struct {
	tUserInfoTable

	EXCLUDED tUserInfoTable
}

// AS creates new TUserInfoTable with assigned alias
func (a TUserInfoTable) AS(alias string) *TUserInfoTable {
	return newTUserInfoTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TUserInfoTable with assigned schema name
func (a TUserInfoTable) FromSchema(schemaName string) *TUserInfoTable {
	return newTUserInfoTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TUserInfoTable with assigned table prefix
func (a TUserInfoTable) WithPrefix(prefix string) *TUserInfoTable {
	return newTUserInfoTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TUserInfoTable with assigned table suffix
func (a TUserInfoTable) WithSuffix(suffix string) *TUserInfoTable {
	return newTUserInfoTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTUserInfoTable(schemaName, tableName, alias string) *TUserInfoTable {
	return &TUserInfoTable{
		tUserInfoTable: newTUserInfoTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newTUserInfoTableImpl("", "excluded", ""),
	}
}

func newTUserInfoTableImpl(schemaName, tableName, alias string) tUserInfoTable {
	var (
		UIDColumn      = postgres.IntegerColumn("uid")
		UsernameColumn = postgres.StringColumn("username")
		PasswordColumn = postgres.StringColumn("password")
		EmailColumn    = postgres.StringColumn("email")
		CreatedColumn  = postgres.DateColumn("created")
		allColumns     = postgres.ColumnList{UIDColumn, UsernameColumn, PasswordColumn, EmailColumn, CreatedColumn}
		mutableColumns = postgres.ColumnList{EmailColumn, CreatedColumn}
	)

	return tUserInfoTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UID:      UIDColumn,
		Username: UsernameColumn,
		Password: PasswordColumn,
		Email:    EmailColumn,
		Created:  CreatedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}