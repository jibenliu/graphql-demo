package router

import (
	"github.com/gin-gonic/gin"
	"graphql-demo/controller"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetRouter() {
	Router.GET("/news/:id", controller.Get)
	Router.GET("/news", controller.List)
	Router.POST("/news", controller.Post)
	Router.PUT("/news/:id", controller.Put)
	Router.DELETE("/news", controller.Destroy)

	Router.POST("/graphql", controller.GraphqlHandler())
	Router.GET("/graphql", controller.GraphqlHandler())
}