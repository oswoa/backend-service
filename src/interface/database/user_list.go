package database

import (
	"github.com/oswoa/backend-service/interface/database/model"
)

func (d Database) UserList(whereQuery string, binds []interface{}, orderBy string) []model.Users {
	var result []model.Users
	d.conn.Query(&result, whereQuery, binds, orderBy)
	return result
}
