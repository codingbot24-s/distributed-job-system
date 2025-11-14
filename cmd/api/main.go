package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codingbot24-s/distributed-job-system/internal/broker"
	"github.com/codingbot24-s/distributed-job-system/pkg/config"
	"github.com/codingbot24-s/distributed-job-system/pkg/logger"
)

func main() {
	// laod the config
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// create a redis client
	rc,err := broker.CreateRedisClient(c.Redis)
	if err != nil {
		log.Fatal(err)
	}
	// check redis connection
	message,err := rc.CheckRedisConnection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	logger.LoggerInit()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	fmt.Println("API server running on :8080")
	http.ListenAndServe(":8080", nil)
}
