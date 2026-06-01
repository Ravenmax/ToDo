package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/Ravenmax/ToDo/internal/core/domain"
	core_http_server "github.com/Ravenmax/ToDo/internal/core/transport/http/server"
)

type StattisticsHTTPHandler struct {
	statisticsService StatisticsService
}
type StatisticsService interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(
	statisticsService StatisticsService,
) *StattisticsHTTPHandler {
	return &StattisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}
func (h *StattisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
