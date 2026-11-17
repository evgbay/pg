package common

import "log"

type Template struct {
	Len     int
	Letters bool
	Digits  bool
	Special bool
}

func NewTemplate(passwordLength *int, enableSpecialSymbols *bool) *Template {
	validator := NewPasswordValidator(MinPasswordLength, MaxPasswordLength)
	err := validator.Validate(passwordLength)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Template{
		Len:     ToValue(passwordLength),
		Letters: true,
		Digits:  true,
		Special: ToValue(enableSpecialSymbols),
	}
}
