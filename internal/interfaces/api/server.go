package api

import (
	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
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
	authController        *controllers.AuthController
	authRepository        repositories.AuthRepository
}

func NewServer(
	engine *gin.Engine,
	reservationController *controllers.ReservationController,
	planController *controllers.PlanController,
	authController *controllers.AuthController,
	authRepository repositories.AuthRepository,
) *Server {
	server := &Server{
		engine:                engine,
		reservationController: reservationController,
		planController:        planController,
		authController:        authController,
		authRepository:        authRepository,
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	// Swagger
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	apiV1 := s.engine.Group("/api/v1")

	// 認証ルートの追加
	v1.SetupAuthRoutes(apiV1, s.authController)

	// 認証が必要なルートにミドルウェアを適用
	protected := apiV1.Group("")
	protected.Use(middleware.AuthMiddleware(s.authRepository))
	{
		v1.SetupReservationRoutes(protected, s.reservationController)
		v1.SetupPlanRoutes(protected, s.planController)
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr) //nolint:wrapcheck // エラーをそのまま返す
}
