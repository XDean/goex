package xecho

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestBreakErrorRecover(t *testing.T) {
	r := BreakErrorRecover()
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/test", nil)
	res := httptest.NewRecorder()

	context := e.NewContext(req, res)
	err := r(func(context echo.Context) error {
		panic(BreakError{Actual: errors.New("must no error")})
		MustNoError(errors.New("must no error"))
		return nil
	})(context)
	assert.Error(t, err, "must no error")
}

func TestBreakErrorRecover2(t *testing.T) {
	r := BreakErrorRecover()
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/test", nil)
	res := httptest.NewRecorder()

	context := e.NewContext(req, res)
	assert.Panics(t, func() {
		r(func(context echo.Context) error {
			panic(errors.New("panic"))
			return nil
		})(context)
	})
}
