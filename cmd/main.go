package main

import (
	"github.com/gin-contrib/cors"
	_ "stontactics/docs"
	"time"

	"github.com/gin-gonic/gin"
	route "stontactics/api/route"
	"stontactics/bootstrap"
)

// @title Ston Tactics
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	app := bootstrap.App()

	env := app.Env
	tinkoffClient := app.TinkoffClient

	if env.AppEnv == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	server := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	server.Use(cors.New(corsConfig))

	route.Setup(env, timeout, db, server, tinkoffClient)

	server.Run(":" + env.Port)
}
