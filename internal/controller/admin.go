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

	user, err := h.services.UserInfo(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, usersResponse{
		Data: []model.User{user},
	})
}

func (h *Handler) banUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.BanUser(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.DeleteUser(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}
