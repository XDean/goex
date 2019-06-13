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

func MustBind(c echo.Context, param interface{}) {
	if err := c.Bind(param); err != nil {
		panic(BreakError{Actual: err})
	}
}

func MustValidate(c echo.Context, param interface{}) {
	if err := c.Validate(param); err != nil {
		panic(BreakError{Actual: err})
	}
}

func MustBindAndValidate(c echo.Context, param interface{}) {
	MustBind(c, param)
	MustValidate(c, param)
}

func MustNoError(err error) {
	if err != nil {
		panic(BreakError{Actual: err})
	}
}
