package usecase

import (
	"context"

	"github.com/AntonioMartinezFernandez/hexagonal-go/config"
	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/eventprocessor"
	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/models"
	"github.com/AntonioMartinezFernandez/hexagonal-go/pkg/logger"
)

// EventProcessor UseCase
type eventProcessorUC struct {
	cfg    *config.Config
	logger logger.Logger
}

// EventProcessor UseCase constructor
func NewEventProcessorUseCase(cfg *config.Config, logger logger.Logger) eventprocessor.UseCase {
	return &eventProcessorUC{cfg: cfg, logger: logger}
}

// Validate EventProcessor route
func (u *eventProcessorUC) ValidateProcessor(ctx context.Context) (*models.EventProcessorValidationResponse, error) {
	response := models.EventProcessorValidationResponse{Validated: true}
	return &response, nil
}

// Execute EventProcessor route
func (u *eventProcessorUC) ExecuteProcessor(ctx context.Context) (*models.EventProcessorResponse, error) {
	response := models.EventProcessorResponse{Response: "validator response"}
	return &response, nil
}
