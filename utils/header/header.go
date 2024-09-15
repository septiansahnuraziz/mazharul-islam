package header

import (
	"github.com/gin-gonic/gin"
	"github.com/mazharul-islam/utils"
	"golang.org/x/net/context"
)

const requestHeaderContextKey string = "github.com/mazharul-islam/utils/header.RequestHeaderInfo"

type RequestHeaderInfo struct {
	RequestID string `json:"request_id"` // same as traceID set on golang context
	Path      string `json:"path"`
	ClientIP  string `json:"client_ip"`
	UserAgent string `json:"user_agent"`
	Referer   string `json:"referer"`
}

func SetRequestHeaderInfoToContext(ginContext *gin.Context, requestHeaderInfo RequestHeaderInfo) {
	ginContext.Set(requestHeaderContextKey, requestHeaderInfo)
}

func GetRequestHeaderInfoFromContext(ctx context.Context) *RequestHeaderInfo {
	requestInfo, ok := ctx.Value(requestHeaderContextKey).(RequestHeaderInfo)
	if !ok {
		return nil
	}

	return &requestInfo
}

func (requestHeaderInfo *RequestHeaderInfo) ToString() string {
	if requestHeaderInfo == nil {
		return ""
	}

	return utils.Dump(requestHeaderInfo)
}
