package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func main(){
	initRedisClient()
	scanOperationFromUser()
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


func scanOperationFromUser() {
	var operation string
	fmt.Println("please enter (s, Set new information), (g, Get information), (d, Delete information:),")
	fmt.Print("	(all, get all keys), (e, Exit): ")
	_, scanErr := fmt.Scanf("%s", &operation)
	showErr(scanErr)

	fmt.Println(operation)
}