package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/do"
	"github.com/samber/mo"
	"koanfexample/config"
	myEcho "koanfexample/echo"
	"koanfexample/f/result"
	"koanfexample/handler"
)

var ih = myEcho.NewInjectorMiddleware(do.New())

func main() {
	do.Provide[config.Config](ih.Get(), config.NewConfig)
	do.Provide[handler.GreetHandler](ih.Get(), handler.NewGreetHandler)
	config.Watch(ih)

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("injector", ih.Get())
			return next(ctx)
		}
	})

	// Routes
	e.GET("/", func(ctx echo.Context) error {
		injector, ok := ctx.Get("injector").(*do.Injector)
		var mi mo.Result[*do.Injector]
		if !ok {
			mi = mo.Err[*do.Injector](errors.New("injector not found"))
		} else {
			mi = mo.Ok(injector)
		}

		return result.Map(mi, func(injector *do.Injector) (handler.GreetHandler, error) {
			return do.Invoke[handler.GreetHandler](injector)
		}).Map(func(h handler.GreetHandler) (handler.GreetHandler, error) {
			return h, h.Greet(ctx)
		}).Error()
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
