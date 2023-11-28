package models

type User struct{
	ID int `json:"id"`
	NAME  string `json:"name"`
	EMAIL string `json:"email"`
}
