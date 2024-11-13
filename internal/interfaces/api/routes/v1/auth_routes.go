package v1

import (
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup, controller *controllers.AuthController) {
	router.POST("/login", controller.Login)
}
