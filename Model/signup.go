package model

type Signup struct {
	Name     string `json:"name"`
	Email    string `josn:"email"`
	Password string `josn:"password"`
}
