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
func CreateRedisClient(redisUrl string) (*redisClientstruct, error) {
	// initialize a local logger instance (quick usage)

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing redis url: %w", err)
	}

	redisClient := redis.NewClient(opt)
	rc := redisClientstruct{
		client: redisClient,
	}

	fmt.Println("redis client created")
	return &rc, nil
}

func (r *redisClientstruct) CheckRedisConnection() (string, error) {
	pong, err := r.client.Ping(ctx).Result()
	if err != nil {
		return "", fmt.Errorf("error getting ping response: %w", err)
	}

	fmt.Println(pong)
	return pong, nil
}
