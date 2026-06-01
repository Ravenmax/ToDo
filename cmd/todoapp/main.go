package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/Ravenmax/ToDo/internal/config"
	core_logger "github.com/Ravenmax/ToDo/internal/core/logger"
	core_pgx_pool "github.com/Ravenmax/ToDo/internal/core/repository/postgres/pull/pull/pgx"
	core_http_middleware "github.com/Ravenmax/ToDo/internal/core/transport/http/middelware"
	core_http_server "github.com/Ravenmax/ToDo/internal/core/transport/http/server"
	tasks_postgres_repository "github.com/Ravenmax/ToDo/internal/features/tasks/repository/postgres"
	tasks_service "github.com/Ravenmax/ToDo/internal/features/tasks/service"
	tasks_transport_http "github.com/Ravenmax/ToDo/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/Ravenmax/ToDo/internal/features/users/repository/postgres"
	users_service "github.com/Ravenmax/ToDo/internal/features/users/service"
	users_transport_http "github.com/Ravenmax/ToDo/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	cfg := core_config.NewConfigMust()
	time.Local = cfg.TimeZone

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

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("inializing posgtress pull")
	//создаем пул подключений к базе даных
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing features", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUserRepository(pool)
	userService := users_service.NewUsersService(usersRepository)
	users_transport_http := users_transport_http.NewUsersHTTPHandler(userService)
	logger.Debug("initializing features", zap.String("feature", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	task_transport_http := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	// создаем транспорт, роуты и апиверсии для сервера и связываем между собой
	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouterV1 := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(users_transport_http.Routes()...)
	apiVersionRouterV1.RegisterRoutes(task_transport_http.Routes()...)

	// Example of usage middleware on router
	// apiVersionRouterV2 := core_http_server.NewApiVersionRouter(
	// 	core_http_server.ApiVersion2,
	// 	core_http_middleware.Dummy("APIV2 middleware"),
	// )
	// apiVersionRouterV2.RegisterRoutes(users_transport_http.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	//стартуем сервер
	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
