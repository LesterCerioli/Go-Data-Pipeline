package initializers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-pipeline/routers"
	pipelineimpl "go-pipeline/services/implementations/pipelines"
)

// Services holds initialized application services for route registration.
type Services struct {
	PipelineService *pipelineimpl.PipelineService
}

// NewService builds the Fiber app and wires its dependencies.
func NewService(db *gorm.DB) *fiber.App {
	services := &Services{
		PipelineService: pipelineimpl.NewPipelineService(db),
	}

	app := fiber.New(fiber.Config{AppName: "go-pipeline"})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	routers.Register(app, services.PipelineService)
	return app
}
