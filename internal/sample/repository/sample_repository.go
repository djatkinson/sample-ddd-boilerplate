package repository

import (
	"context"
	"ddd-boilerplate/internal/sample/entity"
)

type SampleRepository interface {
	FindSampleByID(ctx context.Context, id int64) (*entity.SampleEntity, error)
}
