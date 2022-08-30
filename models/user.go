package models

type User struct {
	UserID       int64  `json:"uid"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	Email        string `json:"user_email"`
	Created      string `json:"user_created"`
}
