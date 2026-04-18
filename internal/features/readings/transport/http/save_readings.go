package readings_http_transport

import (
	"encoding/json"
	"golang-arduino-server/internal/core/domain"
	core_logger "golang-arduino-server/internal/core/logger"
	core_http_response "golang-arduino-server/internal/core/transport/http/response"
	"net/http"
)

type SaveReadingsRequest struct {
	Gasoline         int `json:"gasoline"`
	GeneratorVoltage int `json:"generator_voltage"`
	NetworkVoltage   int `json:"network_voltage"`
}

type SaveReadingsResponse struct {
	ID               int `json:"id"`
	Gasoline         int `json:"gasoline"`
	GeneratorVoltage int `json:"generator_voltage"`
	NetworkVoltage   int `json:"network_voltage"`
}

func (h *ReadingsHTTPHandler) SaveReadings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewResponseHTTPHandler(logger, w)

	var request SaveReadingsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode json",
		)

		return
	}

	readings := domainFromDTO(request)

	domain, err := h.readingsService.SaveReadings(ctx, readings)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to save readings",
		)

		return
	}

	response := dtoFromDomain(domain)

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto SaveReadingsRequest) domain.Readings {
	return domain.NewUnitializedReadings(
		dto.Gasoline,
		dto.GeneratorVoltage,
		dto.NetworkVoltage,
	)
}

func dtoFromDomain(domain domain.Readings) SaveReadingsResponse {
	return SaveReadingsResponse{
		ID:               domain.ID,
		Gasoline:         domain.Gasoline,
		GeneratorVoltage: domain.GeneratorVoltage,
		NetworkVoltage:   domain.NetworkVoltage,
	}
}
