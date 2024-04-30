package api

import (
    "net/http"

    "github.com/labstack/echo/v4"

    "github.com/Equilibriumz/url-shortener/internal/aws"
    "github.com/Equilibriumz/url-shortener/internal/repository"
)

type Handler struct {
    awsService *aws.Service
}

type Request struct {
    URL string `json:"url"`
}

type Response struct {
    Key string `json:"key"`
}

func NewHandler(service *aws.Service) *Handler {
    return &Handler{
        awsService: service,
    }
}

func (h *Handler) WithServerGroup(group *echo.Group) {
    group.GET("/url/:content", h.GetOne)
    group.POST("/url", h.SaveOne)
}

func (h *Handler) GetOne(c echo.Context) error {
    key := c.Param("content")

    repo := repository.NewRepository(h.awsService.DynamoDB)

    record, err := repo.GetByKey(key)
    if err != nil {
        return err
    }

    return c.Redirect(http.StatusMovedPermanently, record.URLString)
}

func (h *Handler) SaveOne(c echo.Context) error {
    r := new(Request)

    if err := c.Bind(r); err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    if err := validateSaveRequest(r); err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    hashID := generateKey()

    repo := repository.NewRepository(h.awsService.DynamoDB)

    err := repo.InsertItem(&repository.Record{
        HashID:    hashID,
        URLString: r.URL,
    })

    if err != nil {
        return c.String(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, Response{Key: hashID})
}
