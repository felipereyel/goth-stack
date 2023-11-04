package models

import "strings"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var EmptyUser = User{}

func GenerateNameFromEmail(email string) string {
	return strings.Split(email, "@")[0]
}
