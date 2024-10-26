package service

import (
	"github.com/oswoa/backend-service/model"
)

// ユーザ一覧照会
func (s Service) UserList(email string, isDeleted bool) []model.UserDetail {
	// queryの実行
	userList := s.database.UserList(email, isDeleted)

	response := make([]model.UserDetail, 0)
	for _, v := range userList {
		val := model.UserDetail{
			UserId:      v.UserId,
			Email:       v.Email,
			IsAvailable: v.IsAvailable,
		}
		response = append(response, val)
	}

	return response
}
