package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_middleware "github.com/Ravenmax/ToDo/internal/core/transport/http/middelware"
	"go.uber.org/zap"
)

type HTTPserver struct {
	mux        *http.ServeMux
	config     Config
	log        *core_logger.Logger
	middleware []core_http_middleware.Middleware
}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPserver {
	return &HTTPserver{
		mux:        http.NewServeMux(),
		config:     config,
		log:        log,
		middleware: middleware,
	}
}

func (h *HTTPserver) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)
		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}
func (h *HTTPserver) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middleware...) //сначала оборачиваем мультплексор в мидлвейры
	//создаем сервер с полученным адресом и мультиплексером уже обернутым в мидлвейры
	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}
	ch := make(chan error, 1) //канал для возвращения ошибок во время работы сервера
	go func() {
		defer close(ch)
		h.log.Warn("Start HTTP server", zap.String("addr", h.config.Addr))
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and server HTTP: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("Shotdown HTTP server")
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("Shotdown HTTP server: %w", err)
		}
		h.log.Warn("HTTP server stoped")
	}
	return nil
}
