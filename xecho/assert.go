package xecho

import "github.com/labstack/echo/v4"

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
