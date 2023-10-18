package handler

import (
	"ddd-boilerplate/internal/app/service"
	"ddd-boilerplate/pkg/responses"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SampleHandler struct {
	sampleService service.SampleService
}

func NewSampleHandler(sampleService service.SampleService) *SampleHandler {
	return &SampleHandler{
		sampleService: sampleService,
	}
}

func (s *SampleHandler) GetSampleByID(ctx *fiber.Ctx) error {
	id, _ := strconv.ParseInt(ctx.Params("id"), 10, 64)
	sampleResponse, err := s.sampleService.FindSampleByID(ctx.Context(), id)
	if err != nil {
		return ctx.JSON(responses.Error([]string{"error", err.Error()}))
	}

	return ctx.JSON(responses.Success(sampleResponse))
}
