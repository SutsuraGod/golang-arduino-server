package core_http_middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(
	handler http.Handler,
	middleware ...Middleware,
) http.Handler {
	if len(middleware) == 0 {
		return handler
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	return handler
}
