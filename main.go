package main

import (
	"log"

	"github.com/heroku/test/dbaccess/myredis"
	"github.com/heroku/test/dbaccess/mysql"
	"github.com/heroku/test/handler"
)

func main() {
	// db接続
	db, err := mysql.OpenMySQL()
	if err != nil {
		log.Print(err)
		return
	}
	rd, err := myredis.OpenRedis()
	if err != nil {
		log.Print(err)
		return
	}
	in := handler.InitDB{DB: db, RD: rd}
	// linebothandleの呼び出し
	in.Linebothandler()
}
