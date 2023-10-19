package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrWrongIdType  = errors.New("wrong user id type")
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
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set("userId", userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get("userId")
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user not found")
		return 0, ErrUserNotFound
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "wrong user id type")
		return 0, ErrWrongIdType
	}

	return idInt, nil
}
