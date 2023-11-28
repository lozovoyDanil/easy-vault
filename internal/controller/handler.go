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

	api := router.Group("/api" /*, h.userIdentity*/)
	{
		profile := api.Group("/profile")
		{
			profile.GET("/", h.userInfo)
			profile.POST("/", h.editUser)
		}

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
			}

			groups := spaces.Group("/groups")
			{
				groups.POST("/", h.createGroup)
				// groups.PUT("/:group_id")
				// groups.DELETE("/:group_id")
			}

			units := groups.Group("/units")
			{
				units.GET("/", h.groupUnits)
				// units.GET("/:unit_id")
				// units.POST("/")
				// units.PUT("/")
				// units.DELETE("/:unit_id")
			}
		}

		customer := api.Group("/customer")
		{
			reserv := customer.Group("/curr-reservations")
			{
				reserv.GET("/", h.reservedUnits)
				// reserv.GET("/:unit_id")
				// reserv.GET("/:unit_id/details")

				// reserv.POST("/:unit_id/unlock")
				// reserv.POST("/:unit_id/lock")
			}

			customer.POST("/reserve-unit/:unit_id", h.reserveUnit)
			// customer.POST("/extend-reserv/:unit_id")
			// customer.DELETE("/cancel-reserv/:unit_id")
		}
	}
	return router
}
