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

func NewResetPassowordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rr := repository.NewResetPasswordRepository(db, domain.CollectionResetPasswordCode)
	rc := &controller.ResetPassowrdController{
		ResetPassowrdUsecase: usecase.NewResetPasswordUsecase(ur, rr, timeout),
		Env:           env,
	}

	group.POST("/reset/password", rc.CreateResetPasswordCode)
	group.POST("/reset/password/confirm", rc.ConfirmResetCode)
	group.PUT("/reset/password", rc.ResetPasswordToken)
}
