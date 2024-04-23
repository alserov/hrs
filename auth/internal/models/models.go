package models

type User struct {
	UUID     string `json:"uuid"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserInfo struct {
	UUID     string `json:"uuid"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
