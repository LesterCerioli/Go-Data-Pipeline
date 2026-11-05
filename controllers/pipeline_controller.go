package controllers

import (
	"github.com/gofiber/fiber/v2"

	"go-pipeline/handlers"
	pipelinecontract "go-pipeline/services/contracts/pipelines"
)

// PipelineController exposes the pipeline HTTP endpoints through Fiber.
// Request handling is delegated to the shared PipelineHandler.
type PipelineController struct {
	handler *handlers.PipelineHandler
}

func NewPipelineController(service pipelinecontract.Service) *PipelineController {
	return &PipelineController{handler: handlers.NewPipelineHandler(service)}
}

func (controller *PipelineController) List(c *fiber.Ctx) error {
	return controller.handler.List(c)
}

func (controller *PipelineController) Create(c *fiber.Ctx) error {
	return controller.handler.Create(c)
}

func (controller *PipelineController) Get(c *fiber.Ctx) error {
	return controller.handler.Get(c)
}

func (controller *PipelineController) Update(c *fiber.Ctx) error {
	return controller.handler.Update(c)
}

func (controller *PipelineController) Delete(c *fiber.Ctx) error {
	return controller.handler.Delete(c)
}
