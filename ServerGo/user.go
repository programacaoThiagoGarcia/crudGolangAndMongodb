package main

//User model with register
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

//Users return list user
type Users []User
