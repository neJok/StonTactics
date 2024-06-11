package route

import (
	"time"

	"stontactics/api/middleware"
	"stontactics/bootstrap"
	"stontactics/mongo"

	"github.com/gin-gonic/gin"
	"github.com/nikita-vanyasin/tinkoff"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine, tinkoffClient *tinkoff.Client) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSwaggerRouter(publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewSignUpRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("api")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewStrategyRouter(env, timeout, db, protectedRouter)
	NewFolderRouter(env, timeout, db, protectedRouter)
	NewSpreadingRouter(env, timeout, db, protectedRouter)
	NewPaymentRouter(env, timeout, db, protectedRouter, publicRouter, tinkoffClient)
}
