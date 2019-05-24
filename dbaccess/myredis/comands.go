package myredis

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

// SetKeyValue expireを30秒とし値をセット
func SetKeyValue(key, value string, redisDb *redis.Pool) {
	conn := redisDb.Get()
	conn.Do("SETEX", key, 30, value)
}

// GetKey キーを取得
func GetKey(key string, redisDb *redis.Pool) string {
	conn := redisDb.Get()
	strArry, err := redis.Strings(conn.Do("KEYS", key))
	if err != nil {
		log.Print(err)
	}
	return strArry[0]
}

// GetValue バリューを取得
func GetValue(key string, redisDb *redis.Pool) string {
	conn := redisDb.Get()
	rtnStr, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Print(err)
	}
	return rtnStr
}
