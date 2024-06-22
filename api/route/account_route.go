package route

import (
	"time"

	"github.com/neJok/StonTactics/api/controller"
	"github.com/neJok/StonTactics/api/middleware"
	"github.com/neJok/StonTactics/bootstrap"
	"github.com/neJok/StonTactics/domain"
	"github.com/neJok/StonTactics/mongo"
	"github.com/neJok/StonTactics/repository"
	"github.com/neJok/StonTactics/usecase"

	"github.com/gin-gonic/gin"
)

func NewAccountRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	ac := &controller.AccountController{
		AccountUsecase: usecase.NewAccountUsecase(ur, timeout),
		Env:            env,
	}

	group.GET("/account", ac.Fetch)
	group.DELETE("/account", ac.Delete)

	group.Use(middleware.BodySizeMiddleware(10<<20))
	group.PUT("/account", ac.Update)
}
