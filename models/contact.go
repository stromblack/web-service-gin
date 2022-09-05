package models

type Contacts struct {
	ContactID   int32  `json:"ContactID"`
	CustomerID  int32  `json:"CustomerID"`
	ContactName string `json:"ContactName"`
	Phone       string `json:"Phone"`
	Email       string `json:"Email"`
}
