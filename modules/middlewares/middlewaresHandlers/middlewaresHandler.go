package middlewaresHandlers

import (
	_pkgConfig "github.com/MarkTBSS/057_CORS/config"
	_pkgMiddlewaresMiddlewaresUsecases "github.com/MarkTBSS/057_CORS/modules/middlewares/middlewaresUsecases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type IMiddlewaresHandler interface {
	Cors() fiber.Handler
}

type middlewaresHandler struct {
	cfg                _pkgConfig.IConfig
	middlewaresUsecase _pkgMiddlewaresMiddlewaresUsecases.IMiddlewaresUsecase
}

func MiddlewaresHandler(middlewaresUsecase _pkgMiddlewaresMiddlewaresUsecases.IMiddlewaresUsecase, cfg _pkgConfig.IConfig) IMiddlewaresHandler {
	return &middlewaresHandler{
		cfg:                cfg,
		middlewaresUsecase: middlewaresUsecase,
	}
}

func (h *middlewaresHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}
