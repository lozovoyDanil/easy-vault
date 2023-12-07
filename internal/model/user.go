package model

type User struct {
	Id       int    `db:"id"`
	FullName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `db:"role"`
}

type UserIdentity struct {
	Id   int    `db:"id"`
	Role string `db:"role"`
}

type UpdateUserInput struct {
	Id       int
	FullName *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
