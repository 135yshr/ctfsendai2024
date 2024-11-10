package v1

import (
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPlanRoutes(router *gin.RouterGroup, controller *controllers.PlanController) {
	router.GET("/plans", controller.GetPlans)
}
