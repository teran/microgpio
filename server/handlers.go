package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/teran/microgpio/drivers/raspberrypi/gpio"
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
	return c.JSON(http.StatusOK, &models.EmptyStatus{
		Status: models.StatusOK,
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
			log.Printf("error occurred on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, &models.EmptyStatus{
		Status: models.StatusOK,
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
			log.Printf("error occurred on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, &models.EmptyStatus{
		Status: models.StatusOK,
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
			log.Printf("error occurred on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, &models.EmptyStatus{
		Status: models.StatusOK,
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
			log.Printf("error occurred on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, &models.EmptyStatus{
		Status: models.StatusOK,
	})
}

func (s *Server) value(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a number")
	}

	pin := gpio.New(id)
	value, err := pin.Value()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer func() {
		err := pin.Close()
		if err != nil {
			log.Printf("error occurred on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, &models.StatusWithValue{
		Status: models.StatusOK,
		Value:  value,
	})
}

func (s *Server) mode(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be a number")
	}

	pin := gpio.New(id)
	mode, err := pin.Mode()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer func() {
		err := pin.Close()
		if err != nil {
			log.Printf("error occurred on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, &models.StatusWithMode{
		Status: models.StatusOK,
		Mode:   models.Mode(mode),
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
			log.Printf("error occurred on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, &models.EmptyStatus{
		Status: models.StatusOK,
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
			log.Printf("error occurred on closing Pin object: %s", err)
		}
	}()

	return c.JSON(http.StatusOK, &models.EmptyStatus{
		Status: models.StatusOK,
	})
}
