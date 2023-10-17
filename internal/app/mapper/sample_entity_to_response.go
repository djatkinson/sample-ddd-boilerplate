package mapper

import (
	"ddd-boilerplate/interface/http/response"
	"ddd-boilerplate/internal/sample/entity"
)

func SampleEntityToResponse(sample *entity.Sample) response.SampleResponse {
	return response.SampleResponse{
		ID:   sample.ID,
		Name: sample.Name,
	}
}
