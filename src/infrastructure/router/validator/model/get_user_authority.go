package model

type ValidatorGetUserAuthorityRequest struct {
	UserId string `validate:"required,len=26,alphanum"`
}

type ValidatorGetUserAuthorityResponse struct {
	AuthorityName string `validate:"required,len=2"`
	// TODO: gRPCの仕様でゼロ値の項目は返却されないため、ゼロ値に関わらず返却するようにする
	CanSearch   bool `validate:"required,boolean"`
	CanCreate   bool `validate:"required,boolean"`
	CanUpdate   bool `validate:"required,boolean"`
	CanDelete   bool `validate:"required,boolean"`
	CanApprove  bool `validate:"required,boolean"`
	CanPullBack bool `validate:"required,boolean"`
}
