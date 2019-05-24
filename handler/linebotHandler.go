package handler

import (
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/heroku/test/action"
	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

// HandlingLinebot LINEからのリクエストを受けて応答をハンドリング
func (dbs *DBs) HandlingLinebot() {

	// client生成
	bot, err := linebot.New("0189c809a76170e6c965b62ac5c9f670",
		"hJ5OAGDvemzFZidHYjg1Ihr5SoHs9eqsgUuok/LoW4uXzKD3lEZpqyqDMKti8Q/bp0rb4aVW2zsjFroGMoi5xTZqdWVrGy/CQE/EbozdNI3+Fyvq7sd4O/5EHyFpZ9mMwA7snSk+JzX8WJjNyXUJJAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Print(err)
		return
	}

	configAction := action.ConfigAction{Bot: bot, MySQLDb: dbs.MySQLDb, RedisDb: dbs.RedisDb}

	// // Postのルーティング
	goji.Post("/callback", func(context web.C, writer http.ResponseWriter, request *http.Request) {

		// request取得
		events, err := bot.ParseRequest(request)
		if err == nil {
			// event毎に処理分岐
			for _, event := range events {
				switch event.Type {
				case linebot.EventTypeFollow:
					configAction.ResponseToFollowEvent(event)
				case linebot.EventTypeMessage:
					configAction.ResponseToMessageEvent(event)
				case linebot.EventTypePostback:
					configAction.ResponseToPostBackEvent(event)
				default:
				}
			}
		} else {
			log.Print(err)
		}

	})
	goji.Serve()
}

// DBs ...
type DBs struct {
	MySQLDb *gorm.DB
	RedisDb *redis.Pool
}
