package redisApplication

import (
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func Redis() *redis.Client {
	return redisClient
}
func init() {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		PoolSize: 3,
	})
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err == redis.Nil {
		log.Fatal("Redis异常", err)
	} else if err != nil {
		log.Fatal("失败:", err.Error())
	}
}
