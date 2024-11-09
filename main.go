package main

import (
	"log"

	_ "github.com/135yshr/ctfsendai2024/docs/openapi"
	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	"github.com/135yshr/ctfsendai2024/internal/infrastructure/repositories"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// @title        予約管理システム API
// @version      1.0
// @description  予約管理システムのRESTful API
// @host         localhost:8080
// @BasePath     /api/v1
func main() {
	container := buildContainer()

	if err := container.Invoke(func(server *api.Server) {
		server.Run(":8080")
	}); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	// インフラストラクチャー層
	container.Provide(func() *gin.Engine {
		return gin.Default()
	})
	container.Provide(repositories.NewMemoryReservationRepository)

	// アプリケーション層
	container.Provide(usecases.NewGetUserReservationsUseCase)

	// インターフェース層
	container.Provide(presenters.NewJSONReservationPresenter)
	container.Provide(controllers.NewReservationController)
	container.Provide(api.NewServer)

	return container
}
