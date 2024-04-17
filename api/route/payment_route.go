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
	"github.com/nikita-vanyasin/tinkoff"
)

func NewPaymentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, secureGroup *gin.RouterGroup, defaultGroup *gin.RouterGroup, tinkoffClient *tinkoff.Client) {
	pr := repository.NewPaymentRepository(db, domain.CollectionPayment)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.PaymentController{
		PaymentUsecase: usecase.NewPaymentUsecase(pr, ur, timeout),
		TinkoffClient:	tinkoffClient,
	}

	secureGroup.POST("/payment/create/tinkoff", pc.CreateTinkoff)
	defaultGroup.POST("/payment/callback/tinkoff", pc.CallbackTinkoff)
}
