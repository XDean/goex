package xecho

import (
	"github.com/labstack/echo/v4"
)

type BreakError struct {
	Actual error
}

func BreakErrorRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					e, ok := r.(BreakError)
					if !ok {
						panic(r)
					}
					err = e.Actual
				}
			}()
			return next(c)
		}
	}
}
