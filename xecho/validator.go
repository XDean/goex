package xecho

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err == nil {
		return nil
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
}

func NewValidator() *Validator {
	v := validator.New()
	_ = v.RegisterValidation("regexp", ValidRegexp)
	return &Validator{Validator: v}
}
