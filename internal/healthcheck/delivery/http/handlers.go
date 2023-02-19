package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/AntonioMartinezFernandez/hexagonal-go/config"
	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/healthcheck"
	"github.com/AntonioMartinezFernandez/hexagonal-go/pkg/logger"
	"github.com/AntonioMartinezFernandez/hexagonal-go/pkg/utils"
)

// Healthcheck handlers
type healthcheckHandlers struct {
	cfg    *config.Config
	hcUC   healthcheck.UseCase
	logger logger.Logger
}

// NewHealthcheckHandlers Healthcheck handlers constructor
func NewHealthcheckHandlers(cfg *config.Config, hcUC healthcheck.UseCase, logger logger.Logger) healthcheck.Handlers {
	return &healthcheckHandlers{cfg: cfg, hcUC: hcUC, logger: logger}
}

// Get
// @Summary Get healthcheck
// @Description Get healthcheck
// @Tags Healthcheck
// @Produce  json
// @Success 200 {object} models.Healthcheck
// @Failure 500 {object} httpErrors.RestErr
// @Router /healthcheck [get]
func (h *healthcheckHandlers) Run() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		eventprocessor, _ := h.hcUC.Run(ctx)

		return c.JSON(http.StatusOK, eventprocessor)
	}
}
