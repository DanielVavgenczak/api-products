package dto

type UserInput struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// @dto to login user account
type UserInputLogin struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}
