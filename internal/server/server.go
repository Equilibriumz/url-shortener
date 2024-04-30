package server

import (
    "github.com/Equilibriumz/url-shortener/internal/aws"
    "github.com/labstack/echo/v4"

    "github.com/Equilibriumz/url-shortener/internal/api"
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
