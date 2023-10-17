package service

import (
	"context"
	"ddd-boilerplate/interfaces/http/response"
	"ddd-boilerplate/internal/app/mapper"
	"ddd-boilerplate/internal/sample/repository"
	"ddd-boilerplate/pkg/logger"
	"errors"
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
	log := logger.Ctx(ctx)
	sample, err := s.sampleRepository.FindSampleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if sample == nil {
		return nil, errors.New("data not found")
	}
	log.Info("mantap")

	result := mapper.SampleEntityToResponse(sample)
	return &result, nil
}
