package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func DumpOutGoingContext(c context.Context) string {
	md, _ := metadata.FromOutgoingContext(c)
	return Dump(md)
}

func DumpIncomingContext(c context.Context) string {
	md, _ := metadata.FromIncomingContext(c)
	return Dump(md)
}

func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value("traceID").(string)

	if !ok {
		return ""
	}

	return traceID
}
