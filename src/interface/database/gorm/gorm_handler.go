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
	QueryOne(result interface{}, where string, binds []interface{}, orderBy string)
	Query(result interface{}, whereQuery string, binds []interface{}, orderBy string)
}

type Connection struct {
	conn *gorm.DB
}

// DBコネクションを作成
func NewConnection() *Connection {

	// 環境変数の取得
	host, err := util.GetEnv(config.DB_HOST)
	if err != nil {
		panic(err)
	}

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

	envDbRetryCount, err := util.GetEnv(config.DB_RETRY_COUNT)
	if err != nil {
		panic(err)
	}
	dbRetryCount, _ := strconv.Atoi(envDbRetryCount)

	envTimeout, err := util.GetEnv(config.TIMEOUT)
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
func (c Connection) QueryOne(result interface{}, where string, binds []interface{}, orderBy string) {
	c.conn.Where(where, binds...).Order(orderBy).First(&result)
}

// WHERE検索をする
func (c Connection) Query(result interface{}, whereQuery string, binds []interface{}, orderBy string) {
	c.conn.Where(whereQuery, binds...).Order(orderBy).Find(result)
}
