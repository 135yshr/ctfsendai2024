package main

import (
	"log"

	_ "github.com/135yshr/ctfsendai2024/docs/openapi"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainRepositories "github.com/135yshr/ctfsendai2024/internal/domain/repositories"
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
	container.Provide(func() domainRepositories.ReservationRepository {
		impl := repositories.NewJSONReservationRepositoryImpl("./configs/json/database.json")
		return repositories.NewJSONReservationRepository(impl)
	})
	container.Provide(func() domainRepositories.PlanRepository {
		impl := repositories.NewJSONPlanRepositoryImpl("./configs/json/database.json")
		return repositories.NewJSONPlanRepository(impl)
	})

	// アプリケーション層
	container.Provide(usecases.NewGetUserReservationsUseCase)
	container.Provide(usecases.NewGetPlansUseCase)

	// インターフェース層
	container.Provide(presenters.NewJSONReservationPresenter)
	container.Provide(presenters.NewJSONPlanPresenter)
	container.Provide(controllers.NewReservationController)
	container.Provide(controllers.NewPlanController)
	container.Provide(api.NewServer)

	return container
}
