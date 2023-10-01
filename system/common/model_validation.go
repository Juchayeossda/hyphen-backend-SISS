package common

import (
	"encoding/json"
	"hyphen-backend-SISS/system/exception"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
var allowedExts = []string{".jpg", ".jepg", ".png", ".mp4", ".mov"}

func init() {
	validate = validator.New()

	validate.RegisterValidation("multipart-file-image", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		if fileHeader, ok := field.Interface().(multipart.FileHeader); ok {
			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			for _, allowedExt := range allowedExts {
				if ext == allowedExt {
					return true
				}
			}
		}
		return false
	})

}

func Validate(model any) {

	err := validate.Struct(model)

	if err != nil {
		var messages []map[string]any
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]any{
				"filed":   err.Field(),
				"message": "this filed is " + err.Tag(),
			})
		}

		marshaledMessages, err := json.Marshal(messages)
		exception.PanicLogging(err)

		panic(exception.ValidationError{
			Message: string(marshaledMessages),
		})
	}
}
