package service

import (
	"strings"

	"github.com/oswoa/backend-service/model"
)

func (s Service) UserList(email string, isDeleted bool) []model.UserDetail {
	// queryの生成
	whereQueries := make([]string, 0)
	binds := make([]interface{}, 0)

	whereQueries = append(whereQueries, "email LIKE ?")
	binds = append(binds, "%"+email+"%")

	if !isDeleted {
		whereQueries = append(whereQueries, "is_deleted = ?")
		binds = append(binds, 0)
	}
	whereQuery := strings.Join(whereQueries, " AND ")

	// orderの作成
	orderBy := "email"

	// queryの実行
	userList := s.database.UserList(whereQuery, binds, orderBy)

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
