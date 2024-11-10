package api

import (
	v1 "github.com/135yshr/ctfsendai2024/internal/interfaces/api/routes/v1"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine                *gin.Engine
	reservationController *controllers.ReservationController
	planController        *controllers.PlanController
}

func NewServer(
	engine *gin.Engine,
	reservationController *controllers.ReservationController,
	planController *controllers.PlanController,
) *Server {
	server := &Server{
		engine:                engine,
		reservationController: reservationController,
		planController:        planController,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Swagger
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	apiV1 := s.engine.Group("/api/v1")
	v1.SetupReservationRoutes(apiV1, s.reservationController)
	v1.SetupPlanRoutes(apiV1, s.planController)
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr)
}
