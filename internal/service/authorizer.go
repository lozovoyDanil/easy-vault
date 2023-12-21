package service

import "main.go/internal/model"

type Authorizer interface {
	AuthorizeAccess(user model.UserIdentity, objId int) error
}

func checkAccess(a Authorizer, user model.UserIdentity, objId int) error {
	return a.AuthorizeAccess(user, objId)
}
