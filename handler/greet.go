package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	"koanfexample/config"
)

type (
	GreetHandler interface {
		Greet(ctx echo.Context) error
	}
	GreetHandlerImpl struct {
		msg string
	}
)

func NewGreetHandler(injector *do.Injector) (GreetHandler, error) {
	cfg, err := do.Invoke[config.Config](injector)
	if err != nil {
		return nil, err
	}
	return &GreetHandlerImpl{
		msg: cfg.Greeting.Message,
	}, nil
}

func (h *GreetHandlerImpl) Greet(ctx echo.Context) error {
	return ctx.String(200, h.msg)
}
