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
	bot, err := linebot.New("", "")
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
