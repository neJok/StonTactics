package route

import (
	"time"

	"stontactics/api/controller"
	"stontactics/bootstrap"
	"stontactics/domain"
	"stontactics/mongo"
	"stontactics/repository"
	"stontactics/usecase"

	"github.com/gin-gonic/gin"
)

func NewAccountRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	ac := &controller.AccountController{
		AccountUsecase: usecase.NewAccountUsecase(ur, timeout),
		Env:           env,
	}

	group.GET("/account", ac.Fetch)
	group.DELETE("/account", ac.DeleteAccount)
}
