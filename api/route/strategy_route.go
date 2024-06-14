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
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := &controller.StrategyController{
		StrategyUsecase: usecase.NewStrategyUsecase(sr, timeout),
		AccountUsecase: usecase.NewAccountUsecase(ur, timeout),
	}
	group.GET("/strategy", sc.FetchMany)
	group.POST("/strategy", sc.Create)
	group.GET("/strategy/:id", sc.FetchOne)
	group.PUT("/strategy/:id", sc.Update)
}
