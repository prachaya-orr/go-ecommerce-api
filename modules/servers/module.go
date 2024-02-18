package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prachaya-or/eCommerce-api/modules/middlewares/middlewaresHandlers"
	"github.com/prachaya-or/eCommerce-api/modules/middlewares/middlewaresRepositories"
	"github.com/prachaya-or/eCommerce-api/modules/middlewares/middlewaresUsecases"
	"github.com/prachaya-or/eCommerce-api/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
	mid    middlewaresHandlers.IMiddlewaresHandlers
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandlers) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
		mid:    mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandlers {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	handler := middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)

	return handler
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)

	m.router.Get("/health", handler.HealthCheck)
}
