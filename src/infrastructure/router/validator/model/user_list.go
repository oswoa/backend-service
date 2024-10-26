package model

type ValidatorUserList struct {
	Email     string `validate:"omitempty,alphanum"`
	IsDeleted bool   `validate:"boolean"`
}
