package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/internal/model"
)

type usersResponse struct {
	Data []model.User `json:"data"`
}

func (h *Handler) allUsers(ctx *gin.Context) {
	users, err := h.services.AllUsers()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, usersResponse{
		Data: users,
	})
}

func (h *Handler) userById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.UserById(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, usersResponse{
		Data: []model.User{*user},
	})
}
