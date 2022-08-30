package models

type Category struct {
	Title       string `json:"tite"`
	Description string `json:"description"`
	ParentId    int    `json:"parentid"`
}
