package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) managerRevenue(ctx *gin.Context) {
	id, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	revenue, err := h.services.ManagerRevenue(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, revenue)
}
