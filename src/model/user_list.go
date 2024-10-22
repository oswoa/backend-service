package model

type UserDetail struct {
	UserId      string
	Email       string
	IsAvailable bool
}

func (u UserDetail) ConvertIsAvailable() string {
	if u.IsAvailable {
		return "1"
	} else {
		return "0"
	}
}
