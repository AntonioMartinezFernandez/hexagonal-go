package healthcheck

import "github.com/labstack/echo/v4"

// Healthcheck HTTP Handlers interface
type Handlers interface {
	Run() echo.HandlerFunc
}
