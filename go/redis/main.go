package main

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

func main() {
	opt, err := redis.ParseURL("redis://@localhost:6379/")
	if err != nil {
		log.Fatalln(err)
	}

	client := redis.NewClient(opt)

	err = client.Set(context.Background(), "hello", "world", 0).Err()
	if err != nil {
		log.Fatalln(err)
	}

	val, err := client.Get(context.Background(), "hello").Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(val)
}
