package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var redisClient *redis.Client

func main() {
	redisClient.Del(redisClient.Context(), "test")
	done := make(chan struct{})
	for i := 0; i < 10000; i++ {
		go func(chan struct{}) {
			redisClient.ZIncrBy(redisClient.Context(), "test", 1, "code")
			done <- struct{}{}

		}(done)
	}

	for i := 0; i < 10000; i++ {
		<-done
	}

	fmt.Println(redisClient.ZScore(redisClient.Context(), "test", "code").Val())

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
