package business_error

import (
	"fmt"
)

// リクエストバリデーションエラー
type RequestValidateError struct {
	ErrCode   string
	ErrTag    string
	FieldName string
}

func (e RequestValidateError) Error() string {
	return fmt.Sprintf(ERR_MSG_VALIDATE_REQUEST,
		e.ErrCode,
		e.FieldName,
		e.GetErrDetail())
}

func (e RequestValidateError) GetErrDetail() string {
	var errDetail string

	switch e.ErrTag {
	case "alphanum":
		errDetail = "半角英数以外"

	case "len", "max":
		errDetail = "文字列長"

	default:
		panic(fmt.Sprintf("[%s]: リクエスト形式が不明な文字種です。", e.ErrTag))
	}

	return errDetail
}
