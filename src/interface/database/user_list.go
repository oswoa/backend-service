package database

import (
	"strings"

	"github.com/oswoa/backend-service/interface/database/model"
)

func (d Database) UserList(email string, isDeleted bool) []model.Users {

	whereQueries := make([]string, 0)
	binds := make([]interface{}, 0)

	whereQueries = append(whereQueries, "email LIKE ?")
	binds = append(binds, "%"+email+"%")

	whereQueries = append(whereQueries, "is_deleted = ?")
	binds = append(binds, isDeleted)

	var result []model.Users
	whereQuery := strings.Join(whereQueries, " AND ")
	d.conn.Query(&result, whereQuery, binds)

	return result
}
