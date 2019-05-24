package action

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/heroku/test/dbaccess/myredis"
	"github.com/heroku/test/dbaccess/mysql"

	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ResponseToFollowEvent followEventに対して応答
func (configAction *ConfigAction) ResponseToFollowEvent(event *linebot.Event) {
	if _, err := configAction.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("友達追加ありがとうございます。")).Do(); err != nil {
		configAction.replyErrMessage(event, err)
		return
	}
}

// ResponseToMessageEvent messageEventに対して応答
func (configAction *ConfigAction) ResponseToMessageEvent(event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		configAction.responseToTextMessage(message, event)
	case *linebot.ImageMessage:
		configAction.responseToImageMessage(event)
	}
}

// ResponseToPostBackEvent postBackEventに対して応答
func (configAction *ConfigAction) ResponseToPostBackEvent(event *linebot.Event) {

	key := "key1"
	myredis.SetKeyValue(key, "value1", configAction.RedisDb)

	if _, err := configAction.Bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage(
		"this is a confilm template",
		linebot.NewConfirmTemplate(
			"key oa value",
			linebot.NewMessageAction(
				"key", myredis.GetKey(key, configAction.RedisDb),
			),
			linebot.NewMessageAction(
				"value", myredis.GetValue(key, configAction.RedisDb),
			),
		),
	)).Do(); err != nil {
		configAction.replyErrMessage(event, err)
		return
	}
}

// responseToTextMessage textMessageの時に応答
func (configAction *ConfigAction) responseToTextMessage(message *linebot.TextMessage, event *linebot.Event) {
	if message.Text == "情報" {
		userID := event.Source.UserID
		profile, _ := configAction.Bot.GetProfile(userID).Do()
		if _, err := configAction.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("あなたの名前は「"+profile.DisplayName+"」\n"+"あなたのIDは「"+userID+"」です。")).Do(); err != nil {
			configAction.replyErrMessage(event, err)
			return
		}
	} else {
		if _, err := configAction.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ユーザー情報を取得したい場合は「情報」と入力してください。")).Do(); err != nil {
			configAction.replyErrMessage(event, err)
			return
		}
	}
}

// responseToImageMessage imageMessageの時に応答
func (configAction *ConfigAction) responseToImageMessage(event *linebot.Event) {
	m := mysql.Mytable{}
	configAction.MySQLDb.First(&m, "id=?", 1)
	if _, err := configAction.Bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage(
		"this is a buttons template",
		linebot.NewButtonsTemplate(
			"",
			m.Name,
			"message text",
			linebot.NewMessageAction(
				"text",
				"text",
			),
			linebot.NewPostbackAction(
				"post back",
				"post back",
				"",
				""),
		),
	)).Do(); err != nil {
		configAction.replyErrMessage(event, err)
		return
	}
}

// replyErrMessage エラーが発生した旨を返す
func (configAction *ConfigAction) replyErrMessage(event *linebot.Event, err error) {
	log.Print(err)
	if _, err := configAction.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("エラーが発生したため再度メッセージを送信してください。")).Do(); err != nil {
		log.Print(err)
	}
}

// ConfigAction ClientとEventを保持
type ConfigAction struct {
	Bot     *linebot.Client
	MySQLDb *gorm.DB
	RedisDb *redis.Pool
}
