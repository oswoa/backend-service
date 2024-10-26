package model

type ValidatorUserListRequest struct {
	Email     string `validate:"omitempty,alphanum"`
	IsDeleted bool   `validate:"required,boolean"`
}

type ValidatorUserListResponse struct {
	UserId      string `validate:"required,alphanum"`
	Email       string `validate:"required,email"`
	IsAvailable bool   `validate:"required,boolean"`
}
