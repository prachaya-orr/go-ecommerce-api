package middlewaresHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/prachaya-or/eCommerce-api/config"
	"github.com/prachaya-or/eCommerce-api/modules/middlewares/middlewaresUsecases"
)

type IMiddlewaresHandlers interface {
	Cors() fiber.Handler
}

type middlewaresHandlers struct {
	cfg               config.IConfig
	middlewareUsecase middlewaresUsecases.IMiddlewaresUsecase
}

func MiddlewaresHandler(cfg config.IConfig, middlewaresUsecase middlewaresUsecases.IMiddlewaresUsecase) IMiddlewaresHandlers {
	return &middlewaresHandlers{
		cfg:               cfg,
		middlewareUsecase: middlewaresUsecase,
	}
}

func (h *middlewaresHandlers) Cors() fiber.Handler {

	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}
