package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"status": "ok",
	})

}

func (h *Handler) signIn(ctx *gin.Context) {

}
