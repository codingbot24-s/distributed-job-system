package broker

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type redisClientstruct struct {
	client *redis.Client
}

// enque deque acnknowlege retry
// TODO: create redis client and connect with it
func CreateRedisClient(redisUrl string) (*redisClientstruct,error) {
	opt,err := redis.ParseURL(redisUrl)
	fmt.Println(redisUrl)
	if err != nil {
		return nil,fmt.Errorf("error parsing redis url")
	}
	redisClient := redis.NewClient(opt)
	rc := redisClientstruct {
		client: redisClient,
	}
	return &rc,nil
}

// ping redis for connection check
func(r *redisClientstruct) CheckRedisConnection() (string,error) {
	pong,err := r.client.Ping(ctx).Result()
	if err != nil {
		return "",fmt.Errorf("error getting ping repsonse")
	}

	return pong,nil
}
