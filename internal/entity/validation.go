package entity

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"sync"
)

var (
	// validate singleton, it's thread safe and cached the struct validation rules
	validate *validator.Validate
	initOnce sync.Once

	// singleton regex
	urlRgx *regexp.Regexp
)

func init() {
	initOnce.Do(func() {
		validate = validator.New()

		//	please add new validation if you need custom validation below.
		//	_ = Validate.RegisterValidation("identifier_format", ValidateIdentifier)
	})
}
