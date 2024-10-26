package validator

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
	"github.com/oswoa/backend-service/infrastructure/router/validator/model"
	"github.com/oswoa/backend-service/internal_error/business_error"
)

// パラメータのバリデーションを実施する
func Validate(params interface{}) error {

	validate := validator.New(validator.WithRequiredStructEnabled())
	if resultErr := validate.Struct(params); resultErr != nil {

		for _, validateErr := range resultErr.(validator.ValidationErrors) {
			switch params.(type) {
			case model.ValidatorUserListRequest:
				ret := business_error.RequestValidateError{
					ErrCode:   business_error.ERR_CODE_VALIDATE_REQUEST,
					ErrTag:    validateErr.Tag(),
					FieldName: validateErr.Field(),
				}
				return ret

			case model.ValidatorUserListResponse:
				ret := business_error.ResponseValidateError{
					ErrCode:   business_error.ERR_CODE_VALIDATE_RESPONSE,
					ErrTag:    validateErr.Tag(),
					FieldName: validateErr.Field(),
				}
				return ret

			default:
				return fmt.Errorf("不明なエラー")
			}
		}
	}
	return nil
}
