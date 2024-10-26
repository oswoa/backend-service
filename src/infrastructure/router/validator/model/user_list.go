package model

type ValidatorUserListRequest struct {
	Email     string `validate:"omitempty,max=256,alphanum"`
	IsDeleted bool   `validate:"required,boolean"`
}

type ValidatorUserListResponse struct {
	UserId string `validate:"required,len=26,alphanum"`
	Email  string `validate:"omitempty,max=256,alphanum"`
	// TODO: gRPCの仕様でゼロ値の項目は返却されないため、ゼロ値に関わらず返却するようにする
	IsAvailable bool `validate:"required,boolean"`
}
