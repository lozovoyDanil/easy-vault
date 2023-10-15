package controller

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	test := router.Group("/test")
	{
		test.GET("/test1", h.test)
	}
	/*
		api := router.Group("/api")
		{
			spaces := api.Group("/spaces")
			{
				groups := spaces.Group("/groups")
				{
					storages := groups.Group("/storages")
					{
						storages.GET("/")
					}
				}
			}
		}
	*/
	return router
}
