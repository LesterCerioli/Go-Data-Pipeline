package pipelines

import (
	"context"

	"go-pipeline/models"
)

type Service interface {
	Create(context.Context, *models.Pipeline) error
	List(context.Context) ([]models.Pipeline, error)
	Get(context.Context, uint) (*models.Pipeline, error)
	Update(context.Context, *models.Pipeline) error
	Delete(context.Context, uint) error
}
