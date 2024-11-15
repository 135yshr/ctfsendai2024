//nolint:errcheck // エラーは無視する
package main

import (
	"log"

	"github.com/135yshr/ctfsendai2024/internal/application/usecases"
	domainRepositories "github.com/135yshr/ctfsendai2024/internal/domain/repositories"
	"github.com/135yshr/ctfsendai2024/internal/foundation/logger"
	"github.com/135yshr/ctfsendai2024/internal/infrastructure/repositories"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/api/middleware"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/controllers"
	"github.com/135yshr/ctfsendai2024/internal/interfaces/presenters"
	"github.com/gin-contrib/cors"
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

	container.Provide(logger.NewLogger)

	// インフラストラクチャー層
	if err := container.Provide(func(logger *logger.Logger) *gin.Engine {
		r := gin.New()
		r.Use(gin.Recovery())

		// CORSミドルウェアを追加
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{
			"http://localhost:3000",
			"http://localhost:8080",
			"http://ctfweb2024.sectanlab.jp:8080",
		}
		config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
		config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
		r.Use(cors.New(config))

		r.Use(middleware.RequestLogger(logger))
		r.Use(middleware.LoggerMiddleware(logger))
		r.Static("/static", "./web/static")
		r.ContextWithFallback = true

		return r
	}); err != nil {
		log.Fatal("DIコンテナの設定に失敗しました:", err)
	}
	container.Provide(func() domainRepositories.ReservationRepository {
		repo, err := repositories.NewReservationRepository(databasePath)
		if err != nil {
			log.Fatalf("予約リポジトリの作成に失敗しました: %s", err.Error())
		}

		return repo
	})
	container.Provide(func() domainRepositories.PlanRepository {
		repo, err := repositories.NewPlanRepository(databasePath)
		if err != nil {
			log.Fatalf("プランリポジトリの作成に失敗しました: %s", err.Error())
		}

		return repo
	})
	container.Provide(func() domainRepositories.UserRepository {
		repo, err := repositories.NewUserRepository(databasePath)
		if err != nil {
			log.Fatalf("認証リポジトリの作成に失敗しました: %s", err.Error())
		}

		return repo
	})
	container.Provide(func() domainRepositories.AuthRepository {
		return repositories.NewJWTAuthRepository("secret")
	})

	// アプリケーション層
	container.Provide(usecases.NewLoginUseCase)
	container.Provide(usecases.NewGetUserReservationsUseCase)
	container.Provide(usecases.NewCreateReservationUseCase)
	container.Provide(usecases.NewGetPlansUseCase)
	container.Provide(usecases.NewGetSecretQuestionUseCase)
	container.Provide(usecases.NewSecretLoginUseCase)
	container.Provide(usecases.NewGetUserInfoUseCase)
	container.Provide(usecases.NewGetUsersUseCase)

	// インターフェース層
	container.Provide(presenters.NewReservationPresenter)
	container.Provide(presenters.NewPlanPresenter)
	container.Provide(presenters.NewAuthPresenter)
	container.Provide(presenters.NewUserPresenter)
	container.Provide(controllers.NewReservationController)
	container.Provide(controllers.NewPlanController)
	container.Provide(controllers.NewAuthController)
	container.Provide(controllers.NewUserController)
	container.Provide(api.NewServer)

	return container
}
