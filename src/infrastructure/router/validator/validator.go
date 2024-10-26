package validator

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/oswoa/backend-service/internal_error/business_error"
)

// パラメータのバリデーションを実施する
func Validate(params interface{}) *business_error.ValidateError {

	validate := validator.New(validator.WithRequiredStructEnabled())
	if resultErr := validate.Struct(params); resultErr != nil {
		for _, v := range resultErr.(validator.ValidationErrors) {
			err := &business_error.ValidateError{
				ErrTag:    v.Tag(),
				FieldName: v.Field(),
			}
			return err
		}
	}
	return nil
}
