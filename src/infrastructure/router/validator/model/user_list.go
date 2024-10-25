package model

type ValidatorUserList struct {
	Email     string `validate:"required"`
	IsDeleted bool
}
