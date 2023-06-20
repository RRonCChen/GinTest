package controller

import (
	"net/http"

	"github.com/RRonCChen/ginTest/service"
	"github.com/gin-gonic/gin"
)

// @Tags      Hello
// @Success   200
// @Param     name path string name "name" default(Ron)
// @Router    /v1/hello/{name} [get]
func (c *Controller) SayHello(context *gin.Context) {
	name := context.Param("name")
	setResponse(context, http.StatusOK, service.SayHello(name))
}
