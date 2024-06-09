package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"stontactics/api/controller"
	"stontactics/bootstrap"
	"stontactics/domain"
	"stontactics/mongo"
	"stontactics/repository"
	"stontactics/usecase"
)

func NewSingUpRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rr := repository.NewRegisterCodeRepository(db, domain.CollectionRegisterCode)
	sc := &controller.SingUpController{
		SingUpUsecase: usecase.NewSingUpUsecase(ur, rr, timeout),
		Env:          env,
	}

	group.POST("/singup/register", sc.SingUp)
	group.POST("/singup/comfirm", sc.ComfirmCode)
}
