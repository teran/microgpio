package controller

import (
	"errors"

	"github.com/teran/microgpio/drivers"
	"github.com/teran/microgpio/drivers/raspberrypi/gpio"
	"github.com/teran/microgpio/models"
)

var (
	// ErrNotFound specifies an error returned if pin is not registered
	ErrNotFound = errors.New("no such device found")

	// ErrIncompatibleMode returned when pin mode is not out
	ErrIncompatibleMode = errors.New("the pin mode is not compatible with on/off state(probably 'in')")
)

// Controller type
type Controller struct {
	pins map[string]drivers.Driver
}

// New type returns new instance of *Controller
func New(pins map[string]int) (*Controller, error) {
	c := &Controller{}
	for k, v := range pins {
		pin := gpio.New(v)
		err := pin.Export()
		if err != nil {
			return nil, err
		}

		c.pins[k] = pin
	}

	return c, nil
}

// On sets high bit to the pin
func (c *Controller) On(name string) error {
	pin, ok := c.pins[name]
	if !ok {
		return ErrNotFound
	}

	mode, err := pin.Mode()
	if err != nil {
		return err
	}
	if mode != models.ModeOut {
		err = pin.Output()
		if err != nil {
			return err
		}
	}

	return pin.High()
}

// Off sets low bit to the pin
func (c *Controller) Off(name string) error {
	pin, ok := c.pins[name]
	if !ok {
		return ErrNotFound
	}

	mode, err := pin.Mode()
	if err != nil {
		return err
	}
	if mode != models.ModeOut {
		err = pin.Output()
		if err != nil {
			return err
		}
	}

	return pin.Low()
}

// Status returns the pin status
func (c *Controller) Status(name string) (models.ToggleStatus, error) {
	pin, ok := c.pins[name]
	if !ok {
		return models.ToggleStatusUnknown, ErrNotFound
	}

	mode, err := pin.Mode()
	if err != nil {
		return models.ToggleStatusUnknown, err
	}

	if mode != models.ModeOut {
		return models.ToggleStatusUnknown, ErrIncompatibleMode
	}

	value, err := pin.Value()
	if err != nil {
		return models.ToggleStatusUnknown, err
	}

	if value == 1 {
		return models.ToggleStatusOn, nil
	}
	return models.ToggleStatusOff, nil
}
