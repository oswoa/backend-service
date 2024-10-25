package database

import (
	gorm_handler "github.com/oswoa/backend-service/interface/database/gorm"
	"github.com/oswoa/backend-service/interface/database/model"
)

// database層の実装
type IDatabase interface {
	UserList(whereQuery string, binds []interface{}, orderBy string) []model.Users
}

// database層のinterfaceを実装する構造体
type Database struct {
	conn gorm_handler.IConnection
}

// database層の実体を作成
func NewDatabase() *Database {
	database := new(Database)
	database.conn = gorm_handler.NewConnection()
	return database
}
