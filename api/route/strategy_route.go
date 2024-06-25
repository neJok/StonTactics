package route

import (
	"time"

	"github.com/neJok/StonTactics/api/controller"
	"github.com/neJok/StonTactics/bootstrap"
	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/mongo"
	"github.com/neJok/StonTactics/repository"
	"github.com/neJok/StonTactics/usecase"

	"github.com/gin-gonic/gin"
)

func NewStrategyRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewStrategyRepository(db, domain.CollectionStrategies)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := &controller.StrategyController{
		StrategyUsecase: usecase.NewStrategyUsecase(sr, timeout),
		AccountUsecase:  usecase.NewAccountUsecase(ur, timeout),
	}
	group.GET("/strategy", sc.FetchMany)
	group.POST("/strategy", sc.Create)
	group.GET("/strategy/:id", sc.FetchOne)
	group.PUT("/strategy/:id", sc.Update)
	group.DELETE("/strategy/:id", sc.DeleteOne)
}
