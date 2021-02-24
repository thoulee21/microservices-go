package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/gbrayhan/microservices-go/docs"

	"github.com/gbrayhan/microservices-go/controllers"
)

// @title Boilerplate Golang
// @version 1.0
// @description Documentation's Boilerplate Golang
// @termsOfService http://swagger.io/terms/

// @contact.name Alejandro Gabriel Guerrero
// @contact.url http://github.com/gbrayhan
// @contact.email gbrayhan@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host domain.com
// @BasePath /v1
func ApplicationV1Router(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		// Documentation Swagger
		{
			v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		v1.POST("/example", controllers.ExampleAction)
	}
}
