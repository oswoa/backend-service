package model

type ValidatorUserListRequest struct {
	Email     string `validate:"omitempty,alphanum"`
	IsDeleted bool   `validate:"required,boolean"`
}

type ValidatorUserListResponse struct {
	UserId string `validate:"required,alphanum"`
	Email  string `validate:"required,email"`
	// TODO: gRPCの仕様でゼロ値の項目は返却されないため、ゼロ値に関わらず返却するようにする
	IsAvailable bool `validate:"required,boolean"`
}
