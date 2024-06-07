package logger

import (
	"github.com/labstack/echo/v4"
)

type MiddlewareLogger struct{}

func (log MiddlewareLogger) LogRequest(handler echo.HandlerFunc) echo.HandlerFunc {

}
