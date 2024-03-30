package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"stontactics/api/middleware"
	"stontactics/bootstrap"
	"stontactics/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSwaggerRouter(publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("api")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewStrategyRouter(env, timeout, db, protectedRouter)
	NewFolderRouter(env, timeout, db, protectedRouter)
	NewSpreadingRouter(env, timeout, db, protectedRouter)
}
