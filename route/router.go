package route

import (
	"github.com/RRonCChen/ginTest/controller"
	"github.com/gin-gonic/gin"
)

func SetRoute(engine *gin.Engine) {
	controller := controller.NewController()
	v1 := engine.Group("/v1")
	v1.GET("/hello/:name", controller.SayHello)
}
