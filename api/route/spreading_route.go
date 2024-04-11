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

func NewSpreadingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewSpreadingRepository(db, domain.CollectionSpreadouts)
	pr := repository.NewUserRepository(db, domain.CollectionUser)
	sc := &controller.SpreadingController{
		SpreadingUsecase: usecase.NewSpreadingUsecase(sr, timeout),
		ProfileUsecase: usecase.NewProfileUsecase(pr, timeout),
	}
	group.GET("/spreading", sc.FetchMany)
	group.POST("/spreading", sc.Create)
	group.GET("/spreading/:id", sc.FetchOne)
	group.PUT("/spreading/:id", sc.Update)
}
