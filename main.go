package main

import (
	"log"

	"github.com/heroku/test/dbaccess/myredis"
	"github.com/heroku/test/dbaccess/mysql"
	"github.com/heroku/test/handler"
)

func main() {
	// db接続
	mysqlDb, err := mysql.OpenMySQL()
	if err != nil {
		log.Print(err)
		return
	}
	redisDb, err := myredis.OpenRedis()
	if err != nil {
		log.Print(err)
		return
	}
	dbs := handler.DBs{MySQLDb: mysqlDb, RedisDb: redisDb}
	// HandlingLinebotの呼び出し
	dbs.HandlingLinebot()
}
