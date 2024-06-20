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

func NewResetPassowordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rr := repository.NewResetPasswordRepository(db, domain.CollectionResetPasswordCode)
	rc := &controller.ResetPassowrdController{
		ResetPassowrdUsecase: usecase.NewResetPasswordUsecase(ur, rr, timeout),
		Env:                  env,
	}

	group.POST("/reset/password", rc.CreateResetPasswordCode)
	group.POST("/reset/password/confirm", rc.ConfirmResetCode)
	group.PUT("/reset/password", rc.ResetPasswordToken)
}
