package models

import "strings"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	PswdHash string `json:"pswd_hash"`
}

var EmptyUser = User{}

func GenerateNameFromEmail(email string) string {
	return strings.Split(email, "@")[0]
}
