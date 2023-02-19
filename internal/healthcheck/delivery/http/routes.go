package http

import (
	"github.com/labstack/echo/v4"

	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/healthcheck"
	"github.com/AntonioMartinezFernandez/hexagonal-go/internal/middleware"
)

// Map healthcheck routes
func MapHealthcheckRoutes(commGroup *echo.Group, h healthcheck.Handlers, mw *middleware.MiddlewareManager) {
	commGroup.GET("", h.Run())
	commGroup.GET("/", h.Run())
}
