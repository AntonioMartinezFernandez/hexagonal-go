package eventprocessor

import "github.com/labstack/echo/v4"

// Event Processor HTTP Handlers interface
type Handlers interface {
	ValidateProcessor() echo.HandlerFunc
	ExecuteProcessor() echo.HandlerFunc
}
