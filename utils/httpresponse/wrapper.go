package httpresponse

import (
	"github.com/gin-gonic/gin"
	"github.com/mazharul-islam/config"
	"github.com/mazharul-islam/utils"
)

type WrapperDto struct {
	Message   string
	Data      any
	ErrorCode string
}

type HttpResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewHttpResponse() *HttpResponse {
	return new(HttpResponse)
}

func (h *HttpResponse) WithMessage(message string) *HttpResponse {
	h.Message = message
	return h
}

func (h *HttpResponse) WithData(data any) *HttpResponse {
	h.Data = data
	return h
}

type WrapperResponseDTO struct {
	ID      string `json:"id"`
	AppName string `json:"appName"`
	Version string `json:"version"`
	Build   string `json:"build"`
	HttpResponse
}

func (h *HttpResponse) ToWrapperResponseDTO(ctx *gin.Context, httpStatus int) {
	var wrapperResponseDTO WrapperResponseDTO
	wrapperResponseDTO.ID = utils.GetTraceID(ctx)
	wrapperResponseDTO.AppName = config.AppName()
	wrapperResponseDTO.Build = config.AppBuild()
	wrapperResponseDTO.Version = config.AppVersion()
	wrapperResponseDTO.Data = h.Data
	wrapperResponseDTO.Message = h.Message

	ctx.JSON(httpStatus, wrapperResponseDTO)
}
