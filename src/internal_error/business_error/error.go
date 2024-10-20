package business_error

import "fmt"

// エラーメッセージ定義
const (
	VALIDATE_ERROR_MSG = "不正なパラメータが送信されました。"
)

// エラーコード定義
const (
	VALIDATE_ERROR_CODE = "EB0001"
)

type ValidateError struct {
	FieldName string
	ErrCode   string
	ErrMsg    string
}

func (e ValidateError) Error() string {
	return fmt.Sprintf("%s %s [parameter: %s]\n",
		e.ErrCode,
		e.ErrMsg,
		e.FieldName)
}

func PrintError(apiName interface{}, err error) {
	val, ok := apiName.(string)
	if !ok {
		fmt.Println("不明なエラーです")
	}
	fmt.Printf("[%s]: %s", val, err.Error())
}
