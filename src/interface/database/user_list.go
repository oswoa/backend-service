package database

import (
	"strings"

	"github.com/oswoa/backend-service/interface/database/model"
)

// ユーザ一覧照会
func (d Database) UserList(email string, isDeleted bool) []model.Users {

	// queryの生成
	whereQueries := make([]string, 0)
	binds := make([]interface{}, 0)

	if len(email) > 0 {
		whereQueries = append(whereQueries, "email LIKE ?")
		binds = append(binds, "%"+email+"%")
	}

	if !isDeleted {
		whereQueries = append(whereQueries, "is_deleted = ?")
		binds = append(binds, 0)
	}
	whereQuery := strings.Join(whereQueries, " AND ")

	orderBy := "email"

	var result []model.Users
	d.conn.Query(&result, whereQuery, binds, orderBy)

	return result
}
