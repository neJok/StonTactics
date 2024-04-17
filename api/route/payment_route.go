package route

import (
	"time"

	"stontactics/api/controller"
	"stontactics/bootstrap"
	"stontactics/mongo"

	"github.com/gin-gonic/gin"
	"github.com/nikita-vanyasin/tinkoff"
)

func NewPaymentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup, tinkoffClient *tinkoff.Client) {
	pc := &controller.PaymentController{
		TinkoffClient:          tinkoffClient,
	}

	group.POST("/payment/create/tinkoff", pc.CreateTinkoff)
	group.POST("/payment/callback/tinkoff", pc.CallbackTinkoff)
}
