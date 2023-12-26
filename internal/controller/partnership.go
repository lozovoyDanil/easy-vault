package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPart(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	tier, err := strconv.Atoi(ctx.Param("tier"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Partnership.CreatePart(user.Id, tier)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}

func (h *Handler) managerPart(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	part, err := h.services.Partnership.PartByUserId(user.Id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, part)
}

func (h *Handler) updatePart(ctx *gin.Context) {
	user, err := getUserIdentity(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	tier, err := strconv.Atoi(ctx.Param("tier"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Partnership.UpdatePart(user.Id, tier)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
