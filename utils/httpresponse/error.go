package httpresponse

import (
	"context"

	"github.com/mazharul-islam/config"
	"github.com/mazharul-islam/utils"
	"github.com/sirupsen/logrus"
)

type WrapperErrorResponseDTO struct {
	WrapperResponseDTO
	ErrorCode string `json:"errorCode,omitempty"`
}

type HTTPError struct {
	Code      int
	Message   string
	ErrorCode string
	Data      any `json:"data,omitempty"`
}

func NewHTTPError() *HTTPError {
	return new(HTTPError)
}

func (httpError *HTTPError) Error() string {
	if httpError == nil {
		return ""
	}

	return httpError.Message
}

func (httpError *HTTPError) StatusCode() int {
	if httpError == nil {
		return 0
	}

	return httpError.Code
}

func (httpError *HTTPError) GetErrorCode() string {
	if httpError == nil {
		return ""
	}

	return httpError.ErrorCode
}

func (httpError *HTTPError) WithCode(httpStatusCode int) *HTTPError {
	httpError.Code = httpStatusCode
	return httpError
}

func (httpError *HTTPError) WithMessage(err error) *HTTPError {
	httpError.Message = err.Error()
	return httpError
}

func (httpError *HTTPError) ToResponseWithContext(context context.Context) WrapperErrorResponseDTO {
	logrus.WithContext(context).WithField("httpError", utils.Dump(httpError)).Error(httpError.Error())

	var wrapperErrorResponse WrapperErrorResponseDTO

	wrapperErrorResponse.ID = utils.GetTraceID(context)
	wrapperErrorResponse.AppName = config.AppName()
	wrapperErrorResponse.Version = config.AppVersion()
	wrapperErrorResponse.Build = config.AppBuild()

	wrapperErrorResponse.Message = httpError.Message
	wrapperErrorResponse.ErrorCode = httpError.ErrorCode
	wrapperErrorResponse.Data = httpError.Data

	return wrapperErrorResponse
}
