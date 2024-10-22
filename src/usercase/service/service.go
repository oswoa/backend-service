package service

import (
	"github.com/oswoa/backend-service/interface/database"
	"github.com/oswoa/backend-service/model"
)

// service層の実装
type IService interface {
	UserList(email string, isDeleted bool) []model.UserDetail
}

// service層のinterfaceを実装する構造体
type Service struct {
	// database層へアクセスするためのinterface
	database database.IDatabase
}

// service層の実体を作成
func NewService(database *database.Database) *Service {
	service := new(Service)
	service.database = database
	return service
}
