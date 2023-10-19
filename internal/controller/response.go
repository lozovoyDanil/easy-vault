package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type statusResp struct {
	Status string `json:"status"`
}

type errorResp struct {
	ErrorMessage string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, code int, msg string) {
	logrus.Error(msg)
	ctx.AbortWithStatusJSON(code, errorResp{ErrorMessage: msg})
}
