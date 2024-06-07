package server

import (
	"github.com/labstack/echo/v4"

	"github.com/Equilibriumz/url-shortener/internal/api"
	"github.com/Equilibriumz/url-shortener/internal/aws"
)

func InitServer() *echo.Echo {
	server := echo.New()

	initHandler(server)

	return server
}

func initHandler(server *echo.Echo) {
	urlGroup := server.Group("/api")

	awsService := aws.NewService()

	api.NewHandler(awsService).WithServerGroup(urlGroup)
}
