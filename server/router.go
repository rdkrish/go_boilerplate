package server

import (
	"github.com/gin-gonic/gin"
	"go_boilerplate/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	goService := router.Group("goservice")
	{
		version := goService.Group("v1")
		{
			controller := new(controllers.ServiceController)
			version.GET("/get", controller.Get)
			version.POST("/post", controller.Post)
			version.DELETE("/delete", controller.Delete)
		}
	}
	return router
}