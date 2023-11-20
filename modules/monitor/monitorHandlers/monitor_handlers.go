package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prachaya-or/eCommerce-api/config"
	"github.com/prachaya-or/eCommerce-api/modules/entities"
	"github.com/prachaya-or/eCommerce-api/modules/monitor"
)

type IMonitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMonitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
	// return c.Status(fiber.StatusOK).JSON(res)
}
