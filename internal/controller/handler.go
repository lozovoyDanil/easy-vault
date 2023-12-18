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

	api := router.Group("/api")
	api.Use(h.userIdentity)
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
		manager.Use(h.managerAccess)
		{
			spaces := manager.Group("/spaces")
			{
				spaces.GET("/", h.userSpaces)
				spaces.POST("/", h.createSpace)
				spaces.PUT("/:space_id", h.updateSpace)
				spaces.DELETE("/:space_id", h.deleteSpace)
			}

			groups := manager.Group("/groups")
			{
				groups.POST("/", h.createGroup)
				groups.PUT("/:group_id", h.updateGroup)
				groups.DELETE("/:group_id", h.deleteGroup)
			}

			units := manager.Group("/units")
			{
				units.GET("/", h.groupUnits)
				units.GET("/:unit_id", h.unitById)
				units.POST("/", h.createUnit)
				units.PUT("/:unit_id", h.updateUnit)
				units.DELETE("/:unit_id", h.deleteUnit)
			}
		}

		customer := api.Group("/customer")
		customer.Use(h.customerAccess)
		{
			reserv := customer.Group("/curr-reservations")
			{
				reserv.GET("/", h.reservedUnits)
				reserv.GET("/:unit_id/details", h.unitDetails)

				reserv.POST("/:unit_id/unlock")
				reserv.POST("/:unit_id/lock")
			}

			customer.POST("/reserve-unit/:unit_id", h.reserveUnit)
			customer.POST("/extend-reserv/:unit_id", h.extendReserv)
			customer.DELETE("/cancel-reserv/:unit_id", h.cancelReserv)
		}

		admin := api.Group("/admin")
		admin.Use(h.adminAccess)
		{
			users := admin.Group("/users")
			{
				users.GET("/", h.allUsers)
				users.GET("/:id", h.userById)
				users.POST("/:id/ban", h.banUser)
				users.DELETE("/:id", h.deleteUser)
			}
		}
	}

	return router
}
