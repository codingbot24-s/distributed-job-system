package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codingbot24-s/distributed-job-system/pkg/config"
	"github.com/codingbot24-s/distributed-job-system/pkg/logger"
)

func main() {
	_, err := config.LoadConfig()
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
