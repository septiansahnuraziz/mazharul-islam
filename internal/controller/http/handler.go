package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mazharul-islam/internal/entity"
	"github.com/mazharul-islam/utils/httpresponse"
)

type Router struct {
	matchService entity.IMatchService
}

func RouteService(
	app *gin.RouterGroup,
	matchService entity.IMatchService,
) {
	router := &Router{
		matchService: matchService,
	}

	router.handlers(app)
}

func (r *Router) handlers(app *gin.RouterGroup) {
	app.GET("/ping", ping)

	apiGroupV1 := app.Group("v1")
	{
		r.initMatchURLRoutes(apiGroupV1)
	}
}

func ping(c *gin.Context) {
	response := httpresponse.NewHttpResponse()
	httpresponse.NoContent(c, response)
}
