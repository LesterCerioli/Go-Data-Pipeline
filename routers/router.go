package routers

import (
	"github.com/gofiber/fiber/v2"

	"go-pipeline/cmd/docs"
	"go-pipeline/controllers"
	pipelinecontract "go-pipeline/services/contracts/pipelines"
)

func Register(app *fiber.App, pipelineService pipelinecontract.Service) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	app.Get("/swagger.json", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.Send(docs.SwaggerJSON)
	})
	app.Get("/swagger.yaml", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "application/yaml; charset=utf-8")
		return c.Send(docs.SwaggerYAML)
	})

	pipelineController := controllers.NewPipelineController(pipelineService)
	pipelines := app.Group("/api/v1/pipelines")
	pipelines.Get("/", pipelineController.List)
	pipelines.Post("/", pipelineController.Create)
	pipelines.Get("/:id", pipelineController.Get)
	pipelines.Put("/:id", pipelineController.Update)
	pipelines.Delete("/:id", pipelineController.Delete)
}
