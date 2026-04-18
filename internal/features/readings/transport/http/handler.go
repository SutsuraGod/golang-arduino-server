package readings_http_transport

import (
	"context"
	"golang-arduino-server/internal/core/domain"
	core_http_server "golang-arduino-server/internal/core/transport/http/server"
	"net/http"
)

type ReadingsHTTPHandler struct {
	readingsService ReadingsService
}

type ReadingsService interface {
	SaveReadings(
		ctx context.Context,
		readings domain.Readings,
	) (domain.Readings, error)
}

func NewReadingsHTTPHandler(service ReadingsService) *ReadingsHTTPHandler {
	return &ReadingsHTTPHandler{
		readingsService: service,
	}
}

func (h *ReadingsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/readings",
			Handler: h.SaveReadings,
		},
	}
}
