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

func NewChangeEmailRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	cr := repository.NewChangeEmailRepository(db, domain.CollectionChangeEmailCode)
	cc := &controller.ChangeEmailController{
		ChangeEmailUsecase: usecase.NewChangeEmailUsecase(ur, cr, timeout),
		Env:           env,
	}

	group.POST("/reset/email", cc.CreateChangeEmailCode)
	group.POST("/reset/email/confirm", cc.ConfirmResetCode)
}
