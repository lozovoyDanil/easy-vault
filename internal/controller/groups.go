package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/internal/model"
)

type groupsResponse struct {
	Data []model.StorageGroup `json:"data"`
}

func (h *Handler) spaceGroups(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid space id")
		return
	}

	groups, err := h.services.SpaceGroups(id, spaceId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, groupsResponse{
		Data: groups,
	})
}

func (h *Handler) groupById(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid space id")
		return
	}

	groupId, err := strconv.Atoi(ctx.Param("group_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid group id")
		return
	}

	group, err := h.services.GroupById(id, spaceId, groupId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, groupsResponse{
		Data: []model.StorageGroup{group},
	})
}

func (h *Handler) createGroup(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid space id")
		return
	}

	var group model.StorageGroup
	err = ctx.BindJSON(&group)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "wrong input: group")
		return
	}

	err = h.services.CreateGroup(id, spaceId, group)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}
