package xecho

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

var ValidateRegexps = map[string]*regexp2.Regexp{}

var DefaultRegexs = struct {
	Username *regexp2.Regexp
	Password *regexp2.Regexp
}{
	Username: regexp2.MustCompile("^(?!_)(?![0-9]+$)[a-zA-Z0-9_]+(?<!_)$", 0),
	Password: regexp2.MustCompile("^(?![0-9]+$)(?![a-zA-Z]+$)(?=[A-Za-z])[0-9A-Za-z]{6,16}$", 0),
}

func ValidRegexp(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {
	case reflect.String:
		str := field.String()
		if str == "" {
			return true
		}
		if regex, find := ValidateRegexps[param]; find {
			isMatch, _ := regex.MatchString(str)
			return isMatch
		} else {
			panic(fmt.Sprintf("Bad regex key %v", str))
		}
	default:
		panic(fmt.Sprintf("Bad field type %T", field.Interface()))
	}
}
