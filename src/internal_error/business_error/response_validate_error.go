package business_error

import (
	"fmt"
)

// レスポンスバリデーションエラー
type ResponseValidateError struct {
	ErrCode   string
	ErrTag    string
	FieldName string
}

func (e ResponseValidateError) Error() string {
	return fmt.Sprintf(ERR_MSG_VALIDATE_RESPONSE,
		e.ErrCode,
		e.FieldName,
		e.GetErrDetail())
}

func (e ResponseValidateError) GetErrDetail() string {
	var errDetail string

	switch e.ErrTag {
	case "alphanum":
		errDetail = "文字種が半角英数以外"
	default:
		panic(fmt.Sprintf("[%s]: レスポンス形式が不明な文字種です。", e.ErrTag))
	}

	return errDetail
}
