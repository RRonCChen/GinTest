package main

import (
	_ "github.com/RRonCChen/ginTest/docs"
	"github.com/RRonCChen/ginTest/route"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title        Ron gin swagger demo
// @version      1.0
// @description  using gin with swagger practice
// @host         localhost:8080
func main() {
	engine := gin.Default()

	route.SetRoute(engine)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	engine.Run(":8080")
}
