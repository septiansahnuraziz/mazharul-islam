package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mazharul-islam/internal/entity"
	"github.com/mazharul-islam/internal/service"
	"github.com/mazharul-islam/utils"
	"github.com/mazharul-islam/utils/httpresponse"
	"net/http"
)

var (
	ErrInternalServerError = httpresponse.NewHTTPError().WithCode(http.StatusInternalServerError).WithMessage(service.ErrInternalServerError)

	errorResponse = map[error]*httpresponse.HTTPError{
		service.ErrNotFound:            httpresponse.NewHTTPError().WithCode(http.StatusNotFound).WithMessage(service.ErrNotFound),
		service.ErrBadRequest:          httpresponse.NewHTTPError().WithCode(http.StatusBadRequest).WithMessage(service.ErrBadRequest),
		service.ErrInternalServerError: ErrInternalServerError,
	}

	successResponse = map[string]string{
		"GetListCustomers": "Success Get List Customers",
	}
)

type (
	metaInfo struct {
		TotalItems int `json:"totalItems,omitempty"`
		TotalPages int `json:"totalPages,omitempty"`
	}

	paginationResponse[T any] struct {
		Items []T `json:"items"`
		metaInfo
	}

	cursorPaginationResponse[T any] struct {
		Data []T               `json:"data"`
		Meta entity.CursorInfo `json:"meta"`
	}
)

func toResourcePaginationResponse[T any](count, size int, items []T) paginationResponse[T] {
	pagination := paginationResponse[T]{
		Items: items,
	}

	if size > 0 {
		pagination.metaInfo = metaInfo{
			TotalItems: int(count),
			TotalPages: utils.CalculatePages(uint(count), uint(size)),
		}
	}

	return pagination
}

func toCursorPaginationResponse[T any](cursorInfo entity.CursorInfo, items []T) cursorPaginationResponse[T] {
	pagination := cursorPaginationResponse[T]{
		Data: items,
		Meta: cursorInfo,
	}

	return pagination
}

func httpErrorHandler(context *gin.Context, err error) {
	if err == nil {
		return
	}

	var validatorError validator.ValidationErrors

	switch {
	case errors.As(err, &validatorError):
		type validationErrorResponse struct {
			Field string
			Error string
		}

		validationErrors := err.(validator.ValidationErrors)
		errorsData := make([]validationErrorResponse, len(validationErrors))

		for i, fe := range validationErrors {
			errorsData[i] = validationErrorResponse{utils.ToCamelCase(fe.Field()), msgForTag(fe.Tag())}
		}

		httpresponse.Error(context, &httpresponse.HTTPError{
			Code:    http.StatusBadRequest,
			Message: service.ErrBadRequest.Error(),
			Data:    errorsData,
		})

		return
	default:
		httpError, ok := errorResponse[err]
		if !ok {
			httpresponse.Error(context, ErrInternalServerError)
			return
		}

		httpresponse.Error(context, httpError)
		return
	}
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address"
	case "isUrl":
		return "Invalid URL format"
	case "min":
		return "The value must be greater than or equal to the specified minimum"
	case "max":
		return "The value must be less than or equal to the specified maximum"
	case "unique":
		return "The value must be unique"
	}

	return utils.WriteStringTemplate("Validation failed for the '%s' tag", tag)
}
