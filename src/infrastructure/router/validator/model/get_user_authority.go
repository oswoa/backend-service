package model

type ValidatorGetUserAuthorityRequest struct {
	UserId string `validate:"required,alphanum"`
}

type ValidatorGetUserAuthorityResponse struct {
	AuthorityName string `validate:"required"`
	CanSearch     bool   `validate:"required,boolean"`
	CanCreate     bool   `validate:"required,boolean"`
	CanUpdate     bool   `validate:"required,boolean"`
	CanDelete     bool   `validate:"required,boolean"`
	CanApprove    bool   `validate:"required,boolean"`
	CanPullBack   bool   `validate:"required,boolean"`
}
