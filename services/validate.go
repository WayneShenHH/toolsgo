package services

import (
	"fmt"
	"regexp"

	"github.com/WayneShenHH/toolsgo/tools"
	"github.com/go-playground/validator"
)

// ExampleValid custom validate function
func ExampleValid() {
	type TestStruct struct {
		CommaSplit string `validate:"sequence"`
	}
	s := TestStruct{
		CommaSplit: "1,2,3",
	}
	var reg = regexp.MustCompile(`^\d[\d\,]+\d$`)
	var validate = validator.New()
	validate.RegisterValidation("sequence", func(fl validator.FieldLevel) bool {
		return reg.MatchString(fl.Field().String())
	})
	err := validate.Struct(s)
	if err != nil {
		tools.Log(err.Error())
	}
	tools.Log(fmt.Sprint("Testing validate string =", s.CommaSplit), fmt.Sprint("result =", err == nil))
	s.CommaSplit = "test faild.."
	err = validate.Struct(s)
	if err != nil {
		tools.Log(err.Error())
	}
	tools.Log(fmt.Sprint("Testing validate string =", s.CommaSplit), fmt.Sprint("result =", err == nil))
}
