package server

import (
	"net/http"

	"github.com/go-shop/pkg/res"
	"github.com/labstack/echo/v4"
)

type healthCheck struct {
	App       string `json:"app"`
	StatusMsg string `json: "status"`
}

func (s *server) healthCheckService(c echo.Context) error {

	return res.SuccessResponse(c, http.StatusOK, &healthCheck{
		App:       s.cfg.App.Name,
		StatusMsg: "OK",
	})
}
