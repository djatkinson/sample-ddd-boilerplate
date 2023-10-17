package service

import (
	"context"
	"ddd-boilerplate/interfaces/http/response"
	"ddd-boilerplate/internal/sample/repository"
)

type sampleService struct {
	sampleRepository repository.SampleRepository
}

type SampleService interface {
	FindSampleByID(ctx context.Context, id int64) (*response.SampleResponse, error)
}

func NewSampleService(sampleRepository repository.SampleRepository) SampleService {
	return &sampleService{
		sampleRepository: sampleRepository,
	}
}

func (s *sampleService) FindSampleByID(ctx context.Context, id int64) (*response.SampleResponse, error) {
	return nil, nil
}
