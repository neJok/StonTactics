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

func NewSpreadingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewSpreadingRepository(db, domain.CollectionSpreadouts)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := &controller.SpreadingController{
		SpreadingUsecase: usecase.NewSpreadingUsecase(sr, timeout),
		AccountUsecase:   usecase.NewAccountUsecase(ur, timeout),
	}
	group.GET("/spreading", sc.FetchMany)
	group.POST("/spreading", sc.Create)
	group.GET("/spreading/:id", sc.FetchOne)
	group.PUT("/spreading/:id", sc.Update)
}
