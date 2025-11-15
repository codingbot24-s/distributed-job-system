package route

import (
	handler "github.com/codingbot24-s/distributed-job-system/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// endpoints//

func StartRouter () {
	app := fiber.New()
	
	app.Post("/api/v1/jobs",handler.EnqueueHandler)



	app.Listen(":8000")
}

