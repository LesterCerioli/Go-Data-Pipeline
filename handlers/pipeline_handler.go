package handlers

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-pipeline/models"
	pipelinecontract "go-pipeline/services/contracts/pipelines"
)

type PipelineHandler struct {
	service pipelinecontract.Service
}

func NewPipelineHandler(service pipelinecontract.Service) *PipelineHandler {
	return &PipelineHandler{service: service}
}

func (handler *PipelineHandler) List(c *fiber.Ctx) error {
	pipelines, err := handler.service.List(c.UserContext())
	if err != nil {
		return respondError(c, fiber.StatusInternalServerError, "could not list pipelines")
	}
	return c.JSON(pipelines)
}

func (handler *PipelineHandler) Create(c *fiber.Ctx) error {
	var pipeline models.Pipeline
	if err := c.BodyParser(&pipeline); err != nil {
		return respondError(c, fiber.StatusBadRequest, "invalid pipeline payload")
	}
	if err := validatePipeline(&pipeline); err != nil {
		return respondError(c, fiber.StatusBadRequest, err.Error())
	}
	if err := handler.service.Create(c.UserContext(), &pipeline); err != nil {
		return respondError(c, fiber.StatusInternalServerError, "could not create pipeline")
	}
	return c.Status(fiber.StatusCreated).JSON(pipeline)
}

func (handler *PipelineHandler) Get(c *fiber.Ctx) error {
	id, err := pipelineID(c)
	if err != nil {
		return err
	}
	pipeline, err := handler.service.Get(c.UserContext(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return respondError(c, fiber.StatusNotFound, "pipeline not found")
	}
	if err != nil {
		return respondError(c, fiber.StatusInternalServerError, "could not retrieve pipeline")
	}
	return c.JSON(pipeline)
}

func (handler *PipelineHandler) Update(c *fiber.Ctx) error {
	id, err := pipelineID(c)
	if err != nil {
		return err
	}
	pipeline, err := handler.service.Get(c.UserContext(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return respondError(c, fiber.StatusNotFound, "pipeline not found")
	}
	if err != nil {
		return respondError(c, fiber.StatusInternalServerError, "could not retrieve pipeline")
	}
	if err := c.BodyParser(pipeline); err != nil {
		return respondError(c, fiber.StatusBadRequest, "invalid pipeline payload")
	}
	pipeline.ID = id
	if err := validatePipeline(pipeline); err != nil {
		return respondError(c, fiber.StatusBadRequest, err.Error())
	}
	if err := handler.service.Update(c.UserContext(), pipeline); err != nil {
		return respondError(c, fiber.StatusInternalServerError, "could not update pipeline")
	}
	return c.JSON(pipeline)
}

func (handler *PipelineHandler) Delete(c *fiber.Ctx) error {
	id, err := pipelineID(c)
	if err != nil {
		return err
	}
	if _, err := handler.service.Get(c.UserContext(), id); errors.Is(err, gorm.ErrRecordNotFound) {
		return respondError(c, fiber.StatusNotFound, "pipeline not found")
	} else if err != nil {
		return respondError(c, fiber.StatusInternalServerError, "could not retrieve pipeline")
	}
	if err := handler.service.Delete(c.UserContext(), id); err != nil {
		return respondError(c, fiber.StatusInternalServerError, "could not delete pipeline")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func validatePipeline(pipeline *models.Pipeline) error {
	pipeline.Name = strings.TrimSpace(pipeline.Name)
	if pipeline.Name == "" {
		return errors.New("name is required")
	}
	if pipeline.Source == nil {
		return errors.New("source configuration is required")
	}
	if pipeline.Destination == nil {
		return errors.New("destination configuration is required")
	}
	if pipeline.Status == "" {
		pipeline.Status = "draft"
	}
	pipeline.Status = strings.ToLower(pipeline.Status)
	switch pipeline.Status {
	case "draft", "active", "paused":
		return nil
	default:
		return errors.New("status must be draft, active, or paused")
	}
}

func pipelineID(c *fiber.Ctx) (uint, error) {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil || id == 0 {
		return 0, fiber.NewError(fiber.StatusBadRequest, "invalid pipeline id")
	}
	return uint(id), nil
}

func respondError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{"error": message})
}
