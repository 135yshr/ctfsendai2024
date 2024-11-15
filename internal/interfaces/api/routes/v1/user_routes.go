package v1

import (
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, controller *controllers.UserController) {
	user := router.Group("/users")
	{
		user.GET("/me", controller.GetMe)
		user.GET("", middleware.RequireAdmin(), controller.GetUsers)
	}
}
