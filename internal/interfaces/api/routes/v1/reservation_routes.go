package v1

import (
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func SetupReservationRoutes(router *gin.RouterGroup, controller *controllers.ReservationController) {
	reservations := router.Group("/reservations")
	{
		reservations.GET("", controller.GetUserReservations)
	}
}
