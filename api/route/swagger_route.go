package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwaggerRouter(group *gin.RouterGroup) {
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
