package echo

import (
	"github.com/b2wdigital/goignite/pkg/server/http/router"
	"github.com/labstack/echo/v4"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c echo.Context) error {

	resp, httpCode := router.Health(c.Request().Context())

	return echo2.JSONResponse(c, httpCode, resp, nil)
}