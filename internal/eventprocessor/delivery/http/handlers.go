package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/AntonioMartinezFernandez/hexagonal-go/config"
	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/eventprocessor"
	"github.com/AntonioMartinezFernandez/hexagonal-go/pkg/logger"
	"github.com/AntonioMartinezFernandez/hexagonal-go/pkg/utils"
)

// Event Processor handlers
type eventProcessorHandlers struct {
	cfg    *config.Config
	epUC   eventprocessor.UseCase
	logger logger.Logger
}

// EventProcessor handlers constructor
func NewEventProcessorHandlers(cfg *config.Config, epUC eventprocessor.UseCase, logger logger.Logger) eventprocessor.Handlers {
	return &eventProcessorHandlers{cfg: cfg, epUC: epUC, logger: logger}
}

// Post
// @Summary Event Processor Endpoint Validator
// @Description Event Processor Endpoint Validator
// @Tags Event Processor
// @Produce  json
// @Success 200 {object} models.EventProcessorValidationResponse
// @Failure 500 {object} httpErrors.RestErr
// @Router /event-processor [post]
func (h *eventProcessorHandlers) ValidateProcessor() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		epvRes, _ := h.epUC.ValidateProcessor(ctx)

		return c.JSON(http.StatusOK, epvRes)
	}
}

// Post
// @Summary Event Processor
// @Description Event Processor
// @Tags Event Processor
// @Produce  json
// @Success 200 {object} models.EventProcessorResponse
// @Failure 500 {object} httpErrors.RestErr
// @Router /event-processor/process [post]
func (h *eventProcessorHandlers) ExecuteProcessor() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		epRes, _ := h.epUC.ExecuteProcessor(ctx)

		return c.JSON(http.StatusOK, epRes)
	}
}
