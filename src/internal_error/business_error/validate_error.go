package business_error

// リクエストバリデーションエラー
type ValidateError struct {
	ErrCode   string
	ErrTag    string
	FieldName string
}
