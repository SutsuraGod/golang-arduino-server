package core_http_server

import (
	"fmt"
	core_http_middleware "golang-arduino-server/internal/core/transport/http/middleware"
	"net/http"
)

type ApiVersion string

const (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
)

type ApiVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	middleware []core_http_middleware.Middleware
}

func NewApiVersionRouter(
	apiVersion ApiVersion,
	middleware ...core_http_middleware.Middleware,
) *ApiVersionRouter {
	return &ApiVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
		middleware: middleware,
	}
}

func (r *ApiVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.WithMiddleware())
	}
}

func (r *ApiVersionRouter) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(
		r,
		r.middleware...,
	)
}
