package app

import (
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis40:6379",
		Password: "",
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		//app.Log.Panic(err)
		panic(err)
	}
	//fmt.Println(pong, err)
}
