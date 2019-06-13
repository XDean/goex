package xecho

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	type TestBean struct {
		Username string `validate:"regexp=USERNAME"`
		Password string `validate:"regexp=PASSWORD"`
	}
	ValidateRegexps["USERNAME"] = DefaultRegexs.Username
	ValidateRegexps["PASSWORD"] = DefaultRegexs.Password

	validator := NewValidator()
	err := validator.Validate(TestBean{})
	assert.NoError(t, err)

	err = validator.Validate(TestBean{
		Username: "valid_name",
		Password: "validpwd123",
	})
	assert.NoError(t, err)

	err = validator.Validate(TestBean{
		Username: "_",
		Password: "123",
	})
	assert.Error(t, err)

	type BadRegex struct {
		Username string `validate:"regexp=wrong"`
	}
	assert.Panics(t, func() {
		_ = validator.Validate(BadRegex{
			Username: "name",
		})
	})

	type BadType struct {
		I int `validate:"regexp=wrong"`
	}
	assert.Panics(t, func() {
		_ = validator.Validate(BadType{
			I: 1,
		})
	})
}
