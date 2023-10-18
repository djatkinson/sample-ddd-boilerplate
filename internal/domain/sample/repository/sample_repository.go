package repository

import (
	"context"
	"ddd-boilerplate/internal/domain/sample/entity"
)

type SampleRepository interface {
	FindSampleByID(ctx context.Context, id int64) (*entity.Sample, error)
}
