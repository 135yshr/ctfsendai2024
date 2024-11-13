//nolint:errcheck // エラーは無視する
package main

import (
	"log"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainRepositories "github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/135yshr/ctfsendai2024/internal/infrastructure/repositories"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	_ "github.com/135yshr/ctfsendai2024/docs/openapi"
)

const databasePath = "./configs/json/database.json"

// main は予約管理システムのサーバーを起動します
// @title        予約管理システム API
// @version      1.0
// @description  予約管理システムのRESTful API
// @host         localhost:8080
// @BasePath     /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Bearer Tokenによる認証
// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT
func main() {
	container := buildContainer()

	if err := container.Invoke(func(server *api.Server) {
		if err := server.Run(":8080"); err != nil {
			log.Fatal("サーバーの起動に失敗しました:", err)
		}
	}); err != nil {
		log.Fatal("サーバーの起動に失敗しました:", err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	// インフラストラクチャー層
	if err := container.Provide(func() *gin.Engine {
		return gin.Default()
	}); err != nil {
		log.Fatal("DIコンテナの設定に失敗しました:", err)
	}
	container.Provide(func() domainRepositories.ReservationRepository {
		return repositories.NewJSONReservationRepository(databasePath)
	})
	container.Provide(func() domainRepositories.PlanRepository {
		return repositories.NewJSONPlanRepository(databasePath)
	})
	container.Provide(func() domainRepositories.AuthRepository {
		repo, err := repositories.NewJWTAuthRepository("secret", databasePath)
		if err != nil {
			panic("認証リポジトリの作成に失敗しました: " + err.Error())
		}

		return repo
	})

	// アプリケーション層
	container.Provide(usecases.NewLoginUseCase)
	container.Provide(usecases.NewGetUserReservationsUseCase)
	container.Provide(usecases.NewGetPlansUseCase)

	// インターフェース層
	container.Provide(presenters.NewJSONReservationPresenter)
	container.Provide(presenters.NewJSONPlanPresenter)
	container.Provide(presenters.NewJSONAuthPresenter)
	container.Provide(controllers.NewReservationController)
	container.Provide(controllers.NewPlanController)
	container.Provide(controllers.NewAuthController)
	container.Provide(api.NewServer)

	return container
}
