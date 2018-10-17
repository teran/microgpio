package server

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/teran/microgpio/models"
)

func (s *Server) index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"/":                  "GET",
		"/ping":              "GET",
		"/gpio/:id/export":   "POST",
		"/gpio/:id/unexport": "POST",
		"/gpio/:id/high":     "POST",
		"/gpio/:id/low":      "POST",
		"/gpio/:id/value":    "GET",
		"/gpio/:id/input":    "POST",
		"/gpio/:id/output":   "POST",
		"/gpio/:id/mode":     "GET",
	})
}

func (s *Server) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.PingResponse{
		Status: models.PingStatusOK,
	})
}

func (s *Server) on(c echo.Context) error {
	pinName := c.Param("name")
	if pinName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "pin name must not be empty")
	}

	err := s.c.On(pinName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	status, err := s.c.Status(pinName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.Status{
		Status: status,
	})
}

func (s *Server) off(c echo.Context) error {
	pinName := c.Param("name")
	if pinName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "pin name must not be empty")
	}

	err := s.c.Off(pinName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	status, err := s.c.Status(pinName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.Status{
		Status: status,
	})
}

func (s *Server) status(c echo.Context) error {
	pinName := c.Param("name")
	if pinName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "pin name must not be empty")
	}

	status, err := s.c.Status(pinName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.Status{
		Status: status,
	})
}
