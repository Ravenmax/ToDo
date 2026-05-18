package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_http_middleware "github.com/Ravenmax/ToDo/internal/core/transport/http/middelware"
	core_http_server "github.com/Ravenmax/ToDo/internal/core/transport/http/server"
	users_transport_http "github.com/Ravenmax/ToDo/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {

	ctx, cancel := signal.NotifyContext( //создаем контексты для обработки системных возовов остановки сервера
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust()) //создаем логгер с новым конфигом
	if err != nil {
		fmt.Println("failed to init application logger")
		os.Exit(1)
	}
	defer logger.Close()
	logger.Debug("starting TODOApp")
	// создаем транспорт, роуты и апиверсии для сервера и связываем между собой
	users_transport_http := users_transport_http.NewUsersHTTPHandler(nil)
	user_routes := users_transport_http.Routes()
	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(user_routes...)
	//создаем http сервер с конфигом и логгером и слоями мидлвеера
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	//регистрируем пути версии api
	httpServer.RegisterAPIRouters(apiVersionRouter)
	//стартуем сервер
	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
