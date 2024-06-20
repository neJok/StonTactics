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
	"github.com/nikita-vanyasin/tinkoff"
)

func NewPaymentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, secureGroup *gin.RouterGroup, defaultGroup *gin.RouterGroup, tinkoffClient *tinkoff.Client) {
	pr := repository.NewPaymentRepository(db, domain.CollectionPayment)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.PaymentController{
		PaymentUsecase: usecase.NewPaymentUsecase(pr, ur, timeout),
		TinkoffClient:  tinkoffClient,
	}

	secureGroup.POST("/payment/create/tinkoff", pc.CreateTinkoff)
	defaultGroup.POST("/payment/callback/tinkoff", pc.CallbackTinkoff)
}
