package controller

import (
	"net/http"

	"github.com/RRonCChen/ginTest/service"
	"github.com/gin-gonic/gin"
)

func (c *Controller) SayHello(context *gin.Context) {
	name := context.Param("name")
	setResponse(context, http.StatusOK, service.SayHello(name))
}
