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

func NewSignUpRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rr := repository.NewRegisterCodeRepository(db, domain.CollectionRegisterCode)
	sc := &controller.SignUpController{
		SignUpUsecase: usecase.NewSignUpUsecase(ur, rr, timeout),
		Env:           env,
	}

	group.POST("/signup/register", sc.SignUp)
	group.POST("/signup/confirm", sc.ConfirmSingUpCode)
}
