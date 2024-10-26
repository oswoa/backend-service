package validator

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/oswoa/backend-service/internal_error/business_error"
)

const (
	VALIDATE_REQUEST = iota
	VALIDATE_RESPONSE
)

// パラメータのバリデーションを実施する
func Validate(params interface{}, kind int) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if resultErr := validate.Struct(params); resultErr != nil {
		for _, v := range resultErr.(validator.ValidationErrors) {
			switch kind {
			case VALIDATE_REQUEST:
				return business_error.RequestValidateError{
					ErrCode:   business_error.ERR_CODE_VALIDATE_REQUEST,
					ErrTag:    v.Tag(),
					FieldName: v.Field(),
				}

			case VALIDATE_RESPONSE:
				return business_error.ResponseValidateError{
					ErrCode:   business_error.ERR_CODE_VALIDATE_RESPONSE,
					ErrTag:    v.Tag(),
					FieldName: v.Field(),
				}
			}
		}
	}
	return nil
}
