package http_server

import (
	"strings"

	"github.com/AntonioMartinezFernandez/hexagonal-go/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	healthcheckHttp "github.com/AntonioMartinezFernandez/hexagonal-go/internal/healthcheck/delivery/http"
	healthcheckUseCase "github.com/AntonioMartinezFernandez/hexagonal-go/internal/healthcheck/usecase"

	eventProcessorHttp "github.com/AntonioMartinezFernandez/hexagonal-go/internal/eventprocessor/delivery/http"
	eventProcessorUseCase "github.com/AntonioMartinezFernandez/hexagonal-go/internal/eventprocessor/usecase"

	apiMiddlewares "github.com/AntonioMartinezFernandez/hexagonal-go/internal/middleware"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories

	// Init useCases
	healthcheckUC := healthcheckUseCase.NewHealthcheckUseCase(s.cfg, s.logger)
	eventProcessorUC := eventProcessorUseCase.NewEventProcessorUseCase(s.cfg, s.logger)

	// Init handlers
	healthcheckHandlers := healthcheckHttp.NewHealthcheckHandlers(s.cfg, healthcheckUC, s.logger)
	eventProcessorHandlers := eventProcessorHttp.NewEventProcessorHandlers(s.cfg, eventProcessorUC, s.logger)

	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)

	e.Use(mw.RequestLoggerMiddleware)

	docs.SwaggerInfo.Title = "Hexagonal Go"
	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "docs")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	v1 := e.Group("/api/v1")

	healthcheck := v1.Group("/healthcheck")
	eventProcessor := v1.Group("/event-processor")

	healthcheckHttp.MapHealthcheckRoutes(healthcheck, healthcheckHandlers, mw)
	eventProcessorHttp.MapEventProcessorRoutes(eventProcessor, eventProcessorHandlers, mw)

	return nil
}
