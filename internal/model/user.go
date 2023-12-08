package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:User,alias:u"`

	Id       int    `bun:"id, pk, autoincrement"`
	FullName string `json:"name" bun:"fullName"`
	Email    string `json:"email" bun:"email"`
	Password string `json:"password" bun:"password"`
	Role     string `json:"role" bun:"role"`
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
