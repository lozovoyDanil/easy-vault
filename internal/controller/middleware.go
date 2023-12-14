package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"

	adminRole    = "admin"
	managerRole  = "manager"
	customerRole = "customer"
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
}

func (h *Handler) adminAccess(ctx *gin.Context) {
	role, err := getUserRole(ctx)
	if errors.Is(err, ErrWrongRoleType) {
		newErrorResponse(ctx, http.StatusInternalServerError, ErrWrongRoleType.Error())
		return
	}
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if role != adminRole {
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

	if role != managerRole {
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

	if role != customerRole {
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
