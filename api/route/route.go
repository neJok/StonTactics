package route

import (
	"time"

	"github.com/neJok/StonTactics/api/middleware"
	"github.com/neJok/StonTactics/bootstrap"
	"github.com/neJok/StonTactics/mongo"

	"github.com/gin-gonic/gin"
	"github.com/nikita-vanyasin/tinkoff"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine, tinkoffClient *tinkoff.Client) {
	publicRouter := gin.Group("")
	// All Public APIs
	publicRouter.Static("/media", "./media")
	NewSwaggerRouter(publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewSignUpRouter(env, timeout, db, publicRouter)
	NewResetPassowordRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("api")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewStrategyRouter(env, timeout, db, protectedRouter)
	NewFolderRouter(env, timeout, db, protectedRouter)
	NewSpreadingRouter(env, timeout, db, protectedRouter)
	NewPaymentRouter(env, timeout, db, protectedRouter, publicRouter, tinkoffClient)
	NewChangeEmailRouter(env, timeout, db, protectedRouter)
	NewAccountRouter(env, timeout, db, protectedRouter)
}
