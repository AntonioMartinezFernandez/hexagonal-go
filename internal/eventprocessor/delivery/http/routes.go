package http

import (
	"github.com/labstack/echo/v4"

	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/eventprocessor"
	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/middleware"
)

// Map eventprocessor routes
func MapEventProcessorRoutes(commGroup *echo.Group, h eventprocessor.Handlers, mw *middleware.MiddlewareManager) {
	commGroup.POST("", h.ValidateProcessor())
	commGroup.POST("/", h.ValidateProcessor())
	commGroup.POST("/process", h.ExecuteProcessor())
}
