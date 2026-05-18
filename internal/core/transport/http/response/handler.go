package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/Ravenmax/ToDo/internal/core/errors"
	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	"go.uber.org/zap"
)

type HttpResponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, rw http.ResponseWriter) *HttpResponseHandler {
	return &HttpResponseHandler{
		log: log,
		rw:  rw,
	}
}
func (h *HttpResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)
	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))

	h.errorResponse(
		statusCode,
		err,
		msg)

}
func (h *HttpResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError

	err := fmt.Errorf("unexpected error: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.errorResponse(
		statusCode,
		err,
		msg)
}

func (h *HttpResponseHandler) errorResponse(
	statusCode int,
	err error,
	msg string,
) {
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("Write HTTP response", zap.Error(err))

	}
}
