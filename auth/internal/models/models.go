package models

type User struct {
	UUID       string `json:"uuid"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IsVerified bool   `yaml:"is_verified"`
}

type UserInfo struct {
	UUID       string `json:"uuid"`
	UserName   string `json:"user_name"`
	IsVerified bool   `yaml:"is_verified"`
}

type VerInfo struct {
	Email      string `json:"email"`
	Code       string `json:"code"`
	IsVerified bool   `json:"is_verified"`
}

type RecoverInfo struct {
	Email    string `json:"email"`
	Code     string `json:"code"`
	Password string `json:"password"`
}
