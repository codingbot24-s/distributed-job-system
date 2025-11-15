package main

import (
	"fmt"
	"log"

	"github.com/codingbot24-s/distributed-job-system/internal/broker"
	route "github.com/codingbot24-s/distributed-job-system/internal/http"
	"github.com/codingbot24-s/distributed-job-system/pkg/config"
)

func main() {
	// laod the config
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	// create a redis client
	rc, err := broker.CreateRedisClient(c.Redis)
	if err != nil {
		log.Fatalf("error creating redis client: %v", err)
	}
	// check redis connection
	message, err := rc.CheckRedisConnection()
	if err != nil {
		log.Fatalf("error checking redis connection: %v", err)
	}
	fmt.Println(message)
	if err != nil {
		log.Fatalf("error creating redis client: %v", err)
	}
	route.StartRouter()
}
