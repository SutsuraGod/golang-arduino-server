package transport

import (
	"encoding/json"
	"golang-arduino-server/internal/features/readings/dto"
	"golang-arduino-server/internal/features/readings/service"
	"net/http"

	"go.uber.org/zap"
)

type HTTPHandlers struct {
	ser    service.ReadingsService
	logger *zap.Logger
}

func NewHTTPHandlers(ser service.ReadingsService, logger *zap.Logger) *HTTPHandlers {
	return &HTTPHandlers{
		ser:    ser,
		logger: logger,
	}
}

func (h *HTTPHandlers) SaveReadings(w http.ResponseWriter, r *http.Request) {
	var readingsDTO dto.ReadingsDTO

	if err := json.NewDecoder(r.Body).Decode(&readingsDTO); err != nil {
		h.logger.Error("Ошибка при переводе json в ReadingsDTO: " + err.Error())

		errDTO := dto.NewErrorDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := h.ser.SaveReadings(&readingsDTO); err != nil {
		errDTO := dto.NewErrorDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(readingsDTO, "", "    ")
	if err != nil {
		h.logger.Error("Проблема при переводе ReadingsDTO в json")
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		h.logger.Error("проблема с отправкой http ответа: " + err.Error())
	}

	h.logger.Info("Запрос обработан успешно")
}
