package model

import "github.com/uptrace/bun"

const (
	AdminRole    = "admin"
	ManagerRole  = "manager"
	CustomerRole = "customer"
)

type User struct {
	bun.BaseModel `bun:"table:User,alias:u"`

	Id       int    `bun:"id,pk,autoincrement"`
	FullName string `json:"name" bun:"fullName"`
	Email    string `json:"email" bun:"email,unique,notnull"`
	Password string `json:"password" bun:"password,notnull"`
	Role     string `json:"role" bun:"role,notnull"`
	IsBanned bool   `json:"isBanned" bun:"isBanned"`
}

type UserIdentity struct {
	Id   int    `db:"id"`
	Role string `db:"role"`
}

type UpdateUserInput struct {
	FullName *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
