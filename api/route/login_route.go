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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}

	group.GET("/auth/:provider", lc.BeginLogin)
	group.GET("/auth/:provider/callback", lc.Callback)
	group.POST("/login", lc.LoginEmail)
}
