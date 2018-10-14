package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (s *Server) index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"/":                "GET",
		"/ping":            "GET",
		"/gpio/:id/high":   "POST",
		"/gpio/:id/low":    "POST",
		"/gpio/:id/output": "POST",
	})
}

func (s *Server) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (s *Server) high(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a number")
	}

	err = s.driver.High(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (s *Server) low(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a number")
	}

	err = s.driver.Low(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (s *Server) output(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a number")
	}

	err = s.driver.Output(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
