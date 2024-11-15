package api

import (
	"fmt"

	"github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/135yshr/ctfsendai2024/internal/foundation/logger"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
	v1 "github.com/135yshr/ctfsendai2024/internal/interfaces/api/routes/v1"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/validators"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine                *gin.Engine
	reservationController *controllers.ReservationController
	planController        *controllers.PlanController
	authController        *controllers.AuthController
	userController        *controllers.UserController
	authRepository        repositories.AuthRepository
	logger                *logger.Logger
}

func NewServer(
	engine *gin.Engine,
	reservationController *controllers.ReservationController,
	planController *controllers.PlanController,
	authController *controllers.AuthController,
	userController *controllers.UserController,
	authRepository repositories.AuthRepository,
	logger *logger.Logger,
) *Server {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validators.RegisterCustomValidations(v)
	}

	server := &Server{
		engine:                engine,
		reservationController: reservationController,
		planController:        planController,
		authController:        authController,
		userController:        userController,
		authRepository:        authRepository,
		logger:                logger,
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
		v1.SetupUserRoutes(protected, s.userController)
	}
}

func (s *Server) Run(addr string) error {
	s.logger.Info("サーバーを起動します", "address", addr)
	if err := s.engine.Run(addr); err != nil {
		return fmt.Errorf("サーバーの起動に失敗しました: %w", err)
	}

	return nil
}
