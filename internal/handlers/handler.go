package handler

import (
	"fmt"
	"time"

	"github.com/codingbot24-s/distributed-job-system/internal/broker"
	jobtype "github.com/codingbot24-s/distributed-job-system/internal/job"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// this will contain all the job handler

func EnqueueHandler(c *fiber.Ctx) error {
	// parse the request json
	job := new(jobtype.Job)
	if err := c.BodyParser(job); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "error parsing request body",
		})
	}
	// validate the required field
	// create a new job instance
	job.JobID = uuid.New().String()
	job.EnqueuedAt = time.Now().Unix()
	// call the brokr enque it will return the jobid or error
	jobId, err := broker.EnqueueToRedis(job)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error ":  err.Error(),
			"message": "error in enque",
		})
	}
	fmt.Println("router workin correctly")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"jobid":      jobId,
		"job status": "enqueued",
		"job type":   job.JobType,
	})
}
