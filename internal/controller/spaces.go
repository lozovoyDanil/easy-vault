package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/internal/model"
)

type spacesResp struct {
	Data []model.Space `json:"data"`
}

func (h *Handler) allSpaces(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
	}

	spaces, err := h.services.UserSpaces(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, spacesResp{
		Data: spaces,
	})
}

func (h *Handler) spaceById(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid param: id")
		return
	}

	space, err := h.services.SpaceById(id, spaceId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, spacesResp{
		Data: []model.Space{space},
	})
}

func (h *Handler) createSpace(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	var input model.Space
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "wrong input: space")
		return
	}

	spaceId, err := h.services.CreateSpace(id, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"id": spaceId,
	})
}

func (h *Handler) updateSpace(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid param: id")
		return
	}

	var space model.UpdateSpaceInput
	err = ctx.BindJSON(&space)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateSpace(id, spaceId, space)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) deleteSpace(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	var spaceId int
	err = ctx.BindJSON(&spaceId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.DeleteSpace(id, spaceId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}
