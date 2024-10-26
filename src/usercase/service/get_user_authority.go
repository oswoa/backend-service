package service

import (
	"github.com/oswoa/backend-service/model"
)

// ユーザ権限取得
func (s Service) GetUserAuthority(userId string) model.UserAuthority {
	// queryの実行
	response := s.database.GetUserAuthority(userId)

	return model.UserAuthority{
		AuthorityName: response.AuthorityName,
		CanSearch:     response.CanSearch,
		CanCreate:     response.CanCreate,
		CanUpdate:     response.CanUpdate,
		CanDelete:     response.CanDelete,
		CanApprove:    response.CanApprove,
		CanPullBack:   response.CanPullBack,
	}
}
