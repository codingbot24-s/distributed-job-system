package broker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/codingbot24-s/distributed-job-system/internal/job"
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

// TODO: must be safe for concurrent call
// enque the job into the redis return the job ID or error  
func (r *redisClientstruct) Enqueue(job *job.Job) error {
	// serialize the job into json convert the golang job struct to json
	_,err := json.Marshal(job)
	if err != nil {
		return fmt.Errorf("error marshaling struct %w",err)
	}
	// push the job to redis

	// return the jobId or error

	return nil
}

