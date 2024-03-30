package route

import (
	"github.com/gin-gonic/gin"
	"stontactics/api/controller"
	"stontactics/bootstrap"
	"stontactics/domain"
	"stontactics/mongo"
	"stontactics/repository"
	"stontactics/usecase"
	"time"
)

func NewStrategyRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewStrategyRepository(db, domain.CollectionStrategies)
	sc := &controller.StrategyController{
		StrategyUsecase: usecase.NewStrategyUsecase(sr, timeout),
	}
	group.GET("/strategy", sc.FetchMany)
	group.POST("/strategy", sc.Create)
	group.GET("/strategy/:id", sc.FetchOne)
	group.PUT("/strategy/:id", sc.Update)
}
