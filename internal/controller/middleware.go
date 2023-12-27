package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"main.go/internal/model"
)

const (
	authHeader = "Authorization"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongIdType   = errors.New("wrong user id type")
	ErrWrongRoleType = errors.New("wrong user role type")
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "wrong header type")
		return
	}

	user, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set("userId", user.Id)
	ctx.Set("userRole", user.Role)
	ctx.Set("userIdentity", user)
}

func (h *Handler) adminAccess(ctx *gin.Context) {
	role, err := getUserRole(ctx)
	if errors.Is(err, ErrWrongRoleType) {
		newErrorResponse(ctx, http.StatusBadRequest, ErrWrongRoleType.Error())
		return
	}
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if role != model.AdminRole {
		newErrorResponse(ctx, http.StatusForbidden, "user is not admin")
		return
	}
}

func (h *Handler) managerAccess(ctx *gin.Context) {
	role, err := getUserRole(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if role != model.ManagerRole && role != model.AdminRole {
		newErrorResponse(ctx, http.StatusForbidden, "user is not manager")
		return
	}
}

func (h *Handler) customerAccess(ctx *gin.Context) {
	role, err := getUserRole(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if role != model.CustomerRole && role != model.AdminRole {
		newErrorResponse(ctx, http.StatusForbidden, "user is not customer")
		return
	}
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get("userId")
	if !ok {
		return 0, ErrUserNotFound
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, ErrWrongIdType
	}

	return idInt, nil
}

func getUserRole(ctx *gin.Context) (string, error) {
	role, ok := ctx.Get("userRole")
	if !ok {
		return "", ErrUserNotFound
	}

	roleStr, ok := role.(string)
	if !ok {
		return "", ErrWrongRoleType
	}

	return roleStr, nil
}

func getUserIdentity(ctx *gin.Context) (*model.UserIdentity, error) {
	id, err := getUserId(ctx)
	if err != nil {
		return nil, err
	}

	role, err := getUserRole(ctx)
	if err != nil {
		return nil, err
	}

	identity := &model.UserIdentity{
		Id:   id,
		Role: role,
	}

	return identity, nil
}
