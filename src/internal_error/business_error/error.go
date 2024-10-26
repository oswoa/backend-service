package business_error

import (
	"context"
	"fmt"
	"log"

	"github.com/oswoa/backend-service/config"
)

// エラーメッセージ定義
const (
	VALIDATE_ERROR_MSG = "不正なパラメータが送信されました。"
)

// エラーコード定義
const (
	VALIDATE_ERROR_CODE = "EB0001"
)

// エラーオブジェクト定義
type ValidateError struct {
	FieldName string
	ErrCode   string
	ErrMsg    string
	ErrTag    string
}

func (e ValidateError) Error() string {
	return fmt.Sprintf("%s %s[parameter: %s][エラー詳細: %s]\n",
		e.ErrCode,
		e.ErrMsg,
		e.FieldName,
		e.GetErrDetail())
}

func (e ValidateError) GetErrDetail() string {

	var errDetail string

	switch e.ErrTag {
	case "alphanum":
		errDetail = "半角英数"
	default:
		panic(fmt.Sprintf("[%s]: 不明な文字種", e.ErrTag))
	}

	return errDetail
}

func PrintError(ctx context.Context, err error) {
	apiName, ok := ctx.Value(config.API_NAME).(string)
	if !ok {
		fmt.Println("不明なエラーです")
		return
	}
	log.Printf("[%s]: %s", apiName, err.Error())
}
