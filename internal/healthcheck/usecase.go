//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package healthcheck

import (
	"context"

	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/models"
)

// Healthcheck use case
type UseCase interface {
	Run(ctx context.Context) (*models.Healthcheck, error)
}
