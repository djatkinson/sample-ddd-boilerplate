package service

import (
	"context"
	"ddd-boilerplate/app/mapper"
	"ddd-boilerplate/domain/sample/repository"
	"ddd-boilerplate/http/response"
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
		log.Error(err.Error())
		return nil, errors.New("data not found")
	}
	log.Info("mantap")
	result := mapper.SampleEntityToResponse(sample)
	return &result, nil
}
