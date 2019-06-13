package xecho

import (
	"github.com/labstack/echo/v4"
)

type J map[string]interface{}

func M(msg string) J {
	return J{
		"message": msg,
	}
}

func BindAndValidate(c echo.Context, param interface{}) {
	if err := c.Bind(param); err != nil {
		panic(BreakError{Actual: err})
	}
	if err := c.Validate(param); err != nil {
		panic(BreakError{Actual: err})
	}
}

func NoError(err error) {
	if err != nil {
		panic(BreakError{Actual: err})
	}
}
