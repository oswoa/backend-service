package validator

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/oswoa/backend-service/internal_error/business_error"
)

// リクエストパラメータのバリデーションを実施する
func RequestValidate(req interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if validateErr := validate.Struct(req); validateErr != nil {
		for _, err := range validateErr.(validator.ValidationErrors) {
			return business_error.ValidateError{
				FieldName: err.Field(),
				ErrCode:   business_error.VALIDATE_ERROR_CODE,
				ErrMsg:    business_error.VALIDATE_ERROR_MSG,
			}
		}
	}

	return nil
}
