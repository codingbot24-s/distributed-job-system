package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// this will contain all the job handler

func EnqueueHandler (c* fiber.Ctx) error {
	// parse the request json 
	// validate the required field
	// create a new job instance
	// call the broker enque it will return the jobid or error 
	fmt.Println("router workin correctly")
	return nil
}
