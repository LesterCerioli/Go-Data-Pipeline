package pipelines

import (
	"context"

	"go-pipeline/models"

	"gorm.io/gorm"
)

type PipelineService struct {
	db *gorm.DB
}

func NewPipelineService(db *gorm.DB) *PipelineService {
	return &PipelineService{db: db}
}

func (s *PipelineService) Create(ctx context.Context, pipeline *models.Pipeline) error {
	return s.db.WithContext(ctx).Create(pipeline).Error
}

func (s *PipelineService) List(ctx context.Context) ([]models.Pipeline, error) {
	var pipelines []models.Pipeline
	err := s.db.WithContext(ctx).Order("id asc").Find(&pipelines).Error
	return pipelines, err
}

func (s *PipelineService) Get(ctx context.Context, id uint) (*models.Pipeline, error) {
	var pipeline models.Pipeline
	if err := s.db.WithContext(ctx).First(&pipeline, id).Error; err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (s *PipelineService) Update(ctx context.Context, pipeline *models.Pipeline) error {
	return s.db.WithContext(ctx).Save(pipeline).Error
}

func (s *PipelineService) Delete(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.Pipeline{}, id).Error
}
