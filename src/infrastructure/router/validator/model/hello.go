package model

type ValidatorHello struct {
	Kind int32 `validate:"required,gte=1,lte=2"`
}
