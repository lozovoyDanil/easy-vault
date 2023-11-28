package model

type User struct {
	Id       int    `db:"id"`
	FullName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string
}

type UpdateUserInput struct {
	Id       int
	FullName *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
