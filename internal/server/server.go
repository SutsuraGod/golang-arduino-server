package server

import (
	"errors"
	"golang-arduino-server/internal/features/readings/transport"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *transport.HTTPHandlers
}

func NewHTTPServer(httpHandlers *transport.HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/readings").Methods("POST").HandlerFunc(s.httpHandlers.SaveReadings)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
