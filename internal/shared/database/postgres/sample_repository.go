package postgres

import (
	"context"
	"ddd-boilerplate/internal/sample/entity"
	"gorm.io/gorm"
)

type SampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{
		db: db,
	}
}

func (s *SampleRepository) FindSampleByID(ctx context.Context, id int64) (*entity.SampleEntity, error) {
	return nil, nil
}
