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

	// router.LoadHTMLGlob("./templates/*")

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{

		spaces := api.Group("/spaces")
		{
			spaces.GET("/", h.allSpaces)
			spaces.GET("/:id", h.spaceById)
			groups := spaces.Group(":id/groups")
			{
				groups.GET("/", h.spaceGroups)
				groups.GET("/:group_id", h.groupById)
			}
		}

		manager := api.Group("/manager")
		{
			spaces := manager.Group("/spaces")
			{
				spaces.GET("/", h.userSpaces)
				spaces.POST("/", h.createSpace)
				spaces.PUT("/:id", h.updateSpace)
				spaces.DELETE("/:id", h.deleteSpace)

				groups := spaces.Group(":id/groups")
				{
					groups.POST("/", h.createGroup)
					groups.PUT("/:group_id")
					groups.DELETE("/:group_id")

					units := groups.Group(":group_id/units")
					{
						units.GET("/")
						units.GET("/:unit_id")
						units.POST("/")
						units.PUT("/")
						units.DELETE("/:unit_id")
					}
				}
			}
		}
		// user := api.Group("/user")
	}
	return router
}
