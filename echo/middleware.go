package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type InjectorMiddleware struct {
	in *do.Injector
}

func NewInjectorMiddleware(in *do.Injector) *InjectorMiddleware {
	return &InjectorMiddleware{in: in}
}

func (m *InjectorMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Set("injector", m.in)
		return next(ctx)
	}
}

func (m *InjectorMiddleware) Get() *do.Injector {
	return m.in
}

func (m *InjectorMiddleware) Apply(f func(*do.Injector)) {
	in := m.in.Clone()
	f(in)
	m.in = in
}
