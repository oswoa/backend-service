package gorm_handler

import (
	"fmt"

	"github.com/oswoa/backend-service/config"
	"github.com/oswoa/backend-service/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IConnection interface {
	QueryOne(result interface{}, where string, binds []interface{})
	Query(result interface{}, whereQuery string, binds []interface{})
}

type Connection struct {
	conn *gorm.DB
}

// DBコネクションを作成
func NewConnection() *Connection {
	database, err := util.GetEnv(config.DATABASE)
	if err != nil {
		panic(err)
	}

	user, err := util.GetEnv(config.DB_USER)
	if err != nil {
		panic(err)
	}

	password, err := util.GetEnv(config.DB_PASSWORD)
	if err != nil {
		panic(err)
	}

	option := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?%s",
		user,
		password,
		database,
		option)

	// TODO: コネクションは明示的に閉じる必要あるのか要確認
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DBの接続に失敗しました")
	}

	connection := new(Connection)
	connection.conn = conn

	return connection
}

// 取得した最初の1件を取得
func (c Connection) QueryOne(result interface{}, where string, binds []interface{}) {
	c.conn.Where(where, binds...).First(&result)
}

// WHERE検索をする
func (c Connection) Query(result interface{}, whereQuery string, binds []interface{}) {
	c.conn.Where(whereQuery, binds...).Find(result)
}
