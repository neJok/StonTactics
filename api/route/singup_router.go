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

func NewSignUpRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rr := repository.NewRegisterCodeRepository(db, domain.CollectionRegisterCode)
	sc := &controller.SignUpController{
		SignUpUsecase: usecase.NewSignUpUsecase(ur, rr, timeout),
		Env:           env,
	}

	group.POST("/signup/register", sc.SignUp)
	group.POST("/signup/comfirm", sc.ComfirmCode)
}
