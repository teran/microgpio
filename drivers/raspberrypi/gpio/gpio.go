package gpio

import (
	"fmt"
	"os"

	"github.com/teran/microgpio/drivers"
)

var _ drivers.Driver = &Pin{}

// Pin type
type Pin struct {
	id int
}

// New returns new *GPIOPin instance
func New(id int) *Pin {
	pin := &Pin{
		id: id,
	}

	return pin
}

// Export exports pin to userspace
func (g *Pin) Export() error {
	fp, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, 0770)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write([]byte(fmt.Sprintf("%d", g.id)))
	return err
}

// Unexport unexports pin from userspace
func (g *Pin) Unexport() error {
	fp, err := os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, 0770)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write([]byte(fmt.Sprintf("%d", g.id)))
	return err
}

// Input sets input mode for the pin
func (g *Pin) Input() error {
	fp, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", g.id), os.O_WRONLY, 0770)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write([]byte("in"))
	return err
}

// Output sets output mode for the pin
func (g *Pin) Output() error {
	fp, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", g.id), os.O_WRONLY, 0770)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write([]byte("out"))
	return err
}

// High sets high bit for GPIO pin
func (g *Pin) High() error {
	fp, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/value", g.id), os.O_WRONLY, 0770)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write([]byte("1"))
	return err
}

// Low sets low bit to GPIO pin
func (g *Pin) Low() error {
	fp, err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/value", g.id), os.O_WRONLY, 0770)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write([]byte("0"))
	return err
}

// Close should destroy all the resources allocated by the *Pin object
func (g *Pin) Close() error {
	return nil
}
