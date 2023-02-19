//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package eventprocessor

import (
	"context"

	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/models"
)

// EventProcessor use case
type UseCase interface {
	ValidateProcessor(ctx context.Context) (*models.EventProcessorValidationResponse, error)
	ExecuteProcessor(ctx context.Context) (*models.EventProcessorResponse, error)
}
