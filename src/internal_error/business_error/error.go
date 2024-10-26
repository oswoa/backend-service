package business_error

import (
	"context"
	"fmt"
	"log"

	"github.com/oswoa/backend-service/config"
)

// エラーメッセージ定義
const (
	ERR_MSG_VALIDATE_REQUEST  = "リクエスト形式が不正です。[ErrCode: %s][param: %s][エラー詳細: %s]"
	ERR_MSG_VALIDATE_RESPONSE = "レスポンス形式が不正です。[ErrCode: %s][param: %s][エラー詳細: %s]"
)

// エラーコード定義
const (
	ERR_CODE_VALIDATE_REQUEST  = "EB0001"
	ERR_CODE_VALIDATE_RESPONSE = "EB0002"
)

// エラーオブジェクト定義
func PrintError(ctx context.Context, err error) {
	apiName, ok := ctx.Value(config.API_NAME).(string)
	if !ok {
		fmt.Println("不明なエラーです")
		return
	}
	log.Printf("[%s]: %s", apiName, err.Error())
}
