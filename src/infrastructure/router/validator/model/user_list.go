package model

type ValidatorUserList struct {
	Email     string `validate:"required,email"`
	IsDeleted bool
}
