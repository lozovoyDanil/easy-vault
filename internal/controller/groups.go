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
	_, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("space_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid space id")
		return
	}

	groups, err := h.services.SpaceGroups(spaceId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, groupsResponse{
		Data: groups,
	})
}

func (h *Handler) groupById(ctx *gin.Context) {
	_, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	groupId, err := strconv.Atoi(ctx.Param("group_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid group id")
		return
	}

	group, err := h.services.GroupById(groupId)
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

	spaceId, err := strconv.Atoi(ctx.Param("space_id"))
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

func (h *Handler) deleteGroup(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// spaceId, err := strconv.Atoi(ctx.Param("space_id"))
	// if err != nil {
	// 	newErrorResponse(ctx, http.StatusBadRequest, "invalid space id")
	// 	return
	// }

	groupId, err := strconv.Atoi(ctx.Param("group_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid group id")
		return
	}

	err = h.services.DeleteGroup(id, groupId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) updateGroup(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	groupId, err := strconv.Atoi(ctx.Param("group_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid group id")
		return
	}
	var group model.UpdateGroupInput
	if err := ctx.BindJSON(&group); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateGroup(id, groupId, group)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}
