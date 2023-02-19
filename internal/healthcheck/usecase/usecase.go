package usecase

import (
	"context"

	"github.com/AntonioMartinezFernandez/hexagonal-go/config"
	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/healthcheck"
	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/models"
	"github.com/AntonioMartinezFernandez/hexagonal-go/pkg/logger"
)

// Healthcheck UseCase
type healthcheckUC struct {
	cfg    *config.Config
	logger logger.Logger
}

// Healthcheck UseCase constructor
func NewHealthcheckUseCase(cfg *config.Config, logger logger.Logger) healthcheck.UseCase {
	return &healthcheckUC{cfg: cfg, logger: logger}
}

// Run Healthcheck
func (u *healthcheckUC) Run(ctx context.Context) (*models.Healthcheck, error) {
	healthcheck := models.Healthcheck{Status: "OK"}
	return &healthcheck, nil
}
