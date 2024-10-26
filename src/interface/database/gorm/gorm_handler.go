package gorm_handler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/oswoa/backend-service/config"
	"github.com/oswoa/backend-service/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IConnection interface {
	QueryOne(result interface{}, where interface{}, binds []interface{}, orderBy interface{}) *gorm.DB
	Query(result interface{}, whereQuery interface{}, binds []interface{}, orderBy interface{}) *gorm.DB
	JoinQuery(result interface{}, joinWhere string, joinBinds []interface{}, whereQuery interface{}, whereBinds []interface{}, orderBy interface{}) *gorm.DB
}

type Connection struct {
	conn *gorm.DB
}

// DBコネクションを作成
func NewConnection() *Connection {

	// 環境変数の取得
	host, err := util.GetEnv(config.ENV_DB_HOST)
	if err != nil {
		panic(err)
	}

	database, err := util.GetEnv(config.ENV_DATABASE)
	if err != nil {
		panic(err)
	}

	user, err := util.GetEnv(config.ENV_DB_USER)
	if err != nil {
		panic(err)
	}

	password, err := util.GetEnv(config.ENV_DB_PASSWORD)
	if err != nil {
		panic(err)
	}

	envDbRetryCount, err := util.GetEnv(config.ENV_DB_RETRY_COUNT)
	if err != nil {
		panic(err)
	}
	dbRetryCount, _ := strconv.Atoi(envDbRetryCount)

	envTimeout, err := util.GetEnv(config.ENV_TIMEOUT)
	if err != nil {
		panic(err)
	}
	timeout, _ := strconv.Atoi(envTimeout)

	option := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?%s",
		user,
		password,
		host,
		database,
		option)

	// DB接続のリトライ処理
	connection := new(Connection)
	for i := 0; i < dbRetryCount; i++ {
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if i-1 == dbRetryCount {
			panic("DB接続のリトライ上限に達しました")
		} else if err != nil {
			log.Printf("[リトライ %d回目]DBの接続に失敗しました", i+1)
			time.Sleep(time.Duration(timeout) * time.Duration(time.Millisecond))
		} else {
			connection.conn = conn
			break
		}
	}

	return connection
}

// 取得した最初の1件を取得
func (c Connection) QueryOne(result interface{}, where interface{}, binds []interface{}, orderBy interface{}) *gorm.DB {
	return c.conn.Where(where, binds...).Order(orderBy).First(&result)
}

// WHERE検索
func (c Connection) Query(result interface{}, whereQuery interface{}, binds []interface{}, orderBy interface{}) *gorm.DB {
	return c.conn.Where(whereQuery, binds...).Order(orderBy).Find(result)
}

// 結合検索
func (c Connection) JoinQuery(result interface{}, joinWhere string, joinBinds []interface{}, whereQuery interface{}, whereBinds []interface{}, orderBy interface{}) *gorm.DB {
	return c.conn.Joins(joinWhere, joinBinds...).Where(whereQuery, whereBinds...).Order(orderBy).Find(result)
}
