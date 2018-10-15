package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/teran/microgpio/drivers/raspberrypi/gpio"
)

func (s *Server) index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"/":                  "GET",
		"/ping":              "GET",
		"/gpio/:id/export":   "POST",
		"/gpio/:id/unexport": "POST",
		"/gpio/:id/high":     "POST",
		"/gpio/:id/low":      "POST",
		"/gpio/:id/input":    "POST",
		"/gpio/:id/output":   "POST",
	})
}

func (s *Server) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (s *Server) export(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a number")
	}

	pin := gpio.New(id)
	err = pin.Export()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer func() {
		err := pin.Close()
		if err != nil {
			log.Printf("error occured on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (s *Server) unexport(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a number")
	}

	pin := gpio.New(id)
	err = pin.Unexport()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer func() {
		err := pin.Close()
		if err != nil {
			log.Printf("error occured on closing Pin object: %s", err)
		}
	}()

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

	pin := gpio.New(id)
	err = pin.High()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer func() {
		err := pin.Close()
		if err != nil {
			log.Printf("error occured on closing Pin object: %s", err)
		}
	}()

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

	pin := gpio.New(id)
	err = pin.Low()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer func() {
		err := pin.Close()
		if err != nil {
			log.Printf("error occured on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (s *Server) input(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a number")
	}

	pin := gpio.New(id)
	err = pin.Input()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer func() {
		err := pin.Close()
		if err != nil {
			log.Printf("error occured on closing Pin object: %s", err)
		}
	}()

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

	pin := gpio.New(id)
	err = pin.Output()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer func() {
		err := pin.Close()
		if err != nil {
			log.Printf("error occured on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
