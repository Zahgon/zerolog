package hlog

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

func FromRequest(r *http.Request) *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func NewHandler(log zerolog.Logger) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func URLHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func MethodHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func RequestHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func RemoteAddrHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func getHost(hostPort string) string { _ = "STUB: not implemented"; return "" }

func RemoteIPHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func UserAgentHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func RefererHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func ProtoHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func HTTPVersionHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

type idKey struct{}

func IDFromRequest(r *http.Request) (id xid.ID, ok bool) {
	_ = "STUB: not implemented"
	return *new(xid.ID), false
}

func IDFromCtx(ctx context.Context) (id xid.ID, ok bool) {
	_ = "STUB: not implemented"
	return *new(xid.ID), false
}

func CtxWithID(ctx context.Context, id xid.ID) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func RequestIDHandler(fieldKey, headerName string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func CustomHeaderHandler(fieldKey, header string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func EtagHandler(fieldKey string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func ResponseHeaderHandler(fieldKey, headerName string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func AccessHandler(f func(r *http.Request, status, size int, duration time.Duration)) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func HostHandler(fieldKey string, trimPort ...bool) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}
