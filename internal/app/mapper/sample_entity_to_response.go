package mapper

import (
	"ddd-boilerplate/http/response"
	"ddd-boilerplate/internal/domain/sample/entity"
)

func SampleEntityToResponse(sample *entity.Sample) response.SampleResponse {
	return response.SampleResponse{
		ID:   sample.ID,
		Name: sample.Name,
	}
}
