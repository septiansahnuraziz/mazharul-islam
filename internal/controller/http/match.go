package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mazharul-islam/internal/entity"
	"github.com/mazharul-islam/utils"
	"github.com/mazharul-islam/utils/httpresponse"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (r *Router) initMatchURLRoutes(app *gin.RouterGroup) {
	customers := app.Group("match")
	{
		customers.GET("/recommendations/user/:id", r.GetRecomendations)
	}
}

// Endpoint Get List Recommendation
//
//	@Summary	Endpoint for get list recommendations
//	@Description
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		Accept			header		string							false	"Example: application/json"
//	@Param		Content-Type	header		string							false	"Example: application/json"
//	@Param		Device-Id		header		string							true	"Example: 5d47eb91-bee9-46b8-9104-aea0f40ef1c3"
//	@Param		Source			header		string							true	"Example: eraspace"
//	@Param		id				path		string								true	"User Id"
//	@Param		request			query		entity.RequestFilterUsers	false	"Query Params"
//	@Success	200				{object}	entity.SwaggerResponseOKDTO{data=cursorPaginationResponse[entity.Users]{data=[]entity.Users},meta=entity.CursorInfo{}}
//	@Failure	400				{object}	entity.SwaggerResponseBadRequestDTO{}			"*Notes: Code data will be return null"
//	@Failure	401				{object}	entity.SwaggerResponseUnauthorizedDTO{}			"*Notes: Code data will be return null"
//	@Failure	500				{object}	entity.SwaggerResponseInternalServerErrorDTO{}	"*Notes: Code data will be return null"
//	@Router		/v1/match/recommendations/user/{id}/ [get]
func (r *Router) GetRecomendations(c *gin.Context) {

	logger := logrus.WithContext(c).WithFields(logrus.Fields{
		"context": utils.DumpIncomingContext(c),
	})

	var searchCriteria entity.RequestFilterUsers
	if err := c.ShouldBindQuery(&searchCriteria); err != nil {
		logger.Error(err)
		httpErrorHandler(c, err)
		return
	}

	searchCriteria.SetDefaultValue()

	customers, cursorInfo, err := r.matchService.GetListRecommendations(c, utils.ExpectedUint(c.Param("id")), searchCriteria)
	if err != nil {
		logger.Error(err)
		httpErrorHandler(c, err)
		return
	}

	resourceShortenURL := toCursorPaginationResponse(cursorInfo, customers)

	httpresponse.NewHttpResponse().
		WithData(resourceShortenURL).
		WithMessage(successResponse["GetListCustomers"]).
		ToWrapperResponseDTO(c, http.StatusOK)
}
