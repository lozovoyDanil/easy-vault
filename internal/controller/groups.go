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

	spaceId, err := strconv.Atoi(ctx.Param("id"))
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
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Query("space_id"))
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

	err = h.services.CreateGroup(*user, spaceId, group)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) deleteGroup(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
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

	err = h.services.DeleteGroup(*user, groupId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) updateGroup(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	groupId, err := strconv.Atoi(ctx.Param("group_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid group id")
		return
	}

	var group model.GroupInput
	if err := ctx.ShouldBindJSON(&group); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if group == (model.GroupInput{}) {
		newErrorResponse(ctx, http.StatusBadRequest, "wrong input: group")
		return
	}

	err = h.services.UpdateGroup(*user, groupId, group)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}
