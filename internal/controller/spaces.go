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
	_, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	var filter model.SpaceFilter
	err = ctx.BindJSON(&filter)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	spaces, err := h.services.AllSpaces(filter)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, spacesResp{
		Data: spaces,
	})
}

func (h *Handler) userSpaces(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
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
	_, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid param: id")
		return
	}

	space, err := h.services.SpaceById(spaceId)
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

	ctx.JSON(http.StatusOK, gin.H{
		"id": spaceId,
	})
}

func (h *Handler) updateSpace(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("space_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid param: id")
		return
	}

	var space model.SpaceInput
	err = ctx.BindJSON(&space)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.UpdateSpace(*user, spaceId, space)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}

func (h *Handler) deleteSpace(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	spaceId, err := strconv.Atoi(ctx.Param("space_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid param: id")
		return
	}

	err = h.services.DeleteSpace(*user, spaceId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResp{
		Status: "OK",
	})
}
