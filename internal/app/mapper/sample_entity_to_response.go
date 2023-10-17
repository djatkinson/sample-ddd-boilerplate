package mapper

import (
	"ddd-boilerplate/interfaces/http/response"
	"ddd-boilerplate/internal/sample/entity"
)

func SampleEntityToResponse(sample *entity.Sample) response.SampleResponse {
	return response.SampleResponse{
		ID:   sample.ID,
		Name: sample.Name,
	}
}
