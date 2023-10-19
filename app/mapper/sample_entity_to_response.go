package mapper

import (
	"ddd-boilerplate/domain/sample/entity"
	"ddd-boilerplate/http/response"
)

func SampleEntityToResponse(sample *entity.Sample) response.SampleResponse {
	return response.SampleResponse{
		ID:   sample.ID,
		Name: sample.Name,
	}
}
