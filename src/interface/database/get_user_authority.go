package database

import (
	"github.com/oswoa/backend-service/interface/database/model"
)

// ユーザ権限取得
func (d Database) GetUserAuthority(userId string) model.Authorities {

	// queryの生成
	joinWhere := "JOIN users ON authorities.authority_id = users.authority_id"
	whereQuery := "users.user_id = ?"
	whereBind := []interface{}{userId}

	var result model.Authorities
	d.conn.JoinQuery(&result, joinWhere, nil, whereQuery, whereBind, nil)

	return result
}
