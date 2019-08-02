package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func main(){
	initRedisClient()
}


/**
initialize redis client for store and restore information from redis database
*/
func initRedisClient() {
	fmt.Println("initializing database...")
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "ZGYKf4a4VZ5v3JbduLu7B3v5mCg2wJ5ATf8CGNgZ",
		DB:0,
	})

	_, err := client.Ping().Result()
	showErr(err)
}