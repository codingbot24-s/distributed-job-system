package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	jobtype "github.com/codingbot24-s/distributed-job-system/internal/job"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// redis client instance
var (
	clientMu       sync.RWMutex
	clientInstance *redis.Client
)

type redisClientstruct struct {
	client *redis.Client
}

// enque deque acnknowlege retry
// TODO: create redis client and connect with it
func CreateRedisClient(redisUrl string) (*redisClientstruct, error) {
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing redis url: %w", err)
	}

	redisClient := redis.NewClient(opt)
	rc := redisClientstruct{
		client: redisClient,
	}

	clientMu.Lock()
	clientInstance = redisClient
	clientMu.Unlock()

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

func GetRedisClient() *redis.Client {
	clientMu.RLock()
	c := clientInstance
	clientMu.RUnlock()
	return c
}

// TODO: must be safe for concurrent call
// enque the job into the redis return the job ID or error
// here we only want to push into redis
func EnqueueToRedis(j *jobtype.Job) (string, error) {
	// marshall will return the bytes
	jsonData, err := json.Marshal(*j)
	if err != nil {
		return "", fmt.Errorf("error marshaling struct %w", err)
	}
	// push the job to redis
	r := GetRedisClient()
	if r == nil {
		return "", fmt.Errorf("redis client not initialized")
	}
	streamName := "Jobs"
	entryID, err := r.XAdd(ctx, &redis.XAddArgs{
		Stream: streamName,
		ID:     "*",
		Values: map[string]interface{}{"job": string(jsonData)},
	}).Result()
	if err != nil {
		return "", fmt.Errorf("error adding job to stream: %w", err)
	}

	fmt.Printf("Added entry to stream '%s' with ID: %s\n", streamName, entryID)

	// return the jobId or error
	return entryID, nil
}
