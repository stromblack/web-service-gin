//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Contacts struct {
	ContactID   int32 `sql:"primary_key"`
	CustomerID  *int32
	ContactName string
	Phone       *string
	Email       *string
}