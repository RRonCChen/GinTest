package main

import (
	"github.com/RRonCChen/ginTest/route"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	route.SetRoute(engine)

	engine.Run(":8080")
}
