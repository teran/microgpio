package gpio

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"

	"github.com/teran/microgpio/drivers"
)

var _ drivers.Driver = (*Pin)(nil)

// Pin type
type Pin struct {
	sync.Mutex

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
	return ioutil.WriteFile(
		"/sys/class/gpio/export",
		[]byte(strconv.Itoa(g.id)),
		os.ModeExclusive,
	)
}

// Unexport unexports pin from userspace
func (g *Pin) Unexport() error {
	return ioutil.WriteFile(
		"/sys/class/gpio/unexport",
		[]byte(strconv.Itoa(g.id)),
		os.ModeExclusive,
	)
}

// Input sets input mode for the pin
func (g *Pin) Input() error {
	return ioutil.WriteFile(
		fmt.Sprintf("/sys/class/gpio/gpio%d/direction", g.id),
		[]byte("in"),
		os.ModeExclusive,
	)
}

// Output sets output mode for the pin
func (g *Pin) Output() error {
	return ioutil.WriteFile(
		fmt.Sprintf("/sys/class/gpio/gpio%d/direction", g.id),
		[]byte("out"),
		os.ModeExclusive,
	)
}

// Mode returns pin mode. Normally values should be on of: in, out
func (g *Pin) Mode() (string, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", g.id))
	if err != nil {
		return "", err
	}
	return string(data), err
}

// High sets high bit for GPIO pin
func (g *Pin) High() error {
	return ioutil.WriteFile(
		fmt.Sprintf("/sys/class/gpio/gpio%d/value", g.id),
		[]byte("1"),
		os.ModeExclusive,
	)
}

// Low sets low bit to GPIO pin
func (g *Pin) Low() error {
	return ioutil.WriteFile(
		fmt.Sprintf("/sys/class/gpio/gpio%d/value", g.id),
		[]byte("0"),
		os.ModeExclusive,
	)
}

// Value returns current value set for the pin
func (g *Pin) Value() (int, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("/sys/class/gpio/gpio%d/value", g.id))
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(data))
}

// Close should destroy all the resources allocated by the *Pin object
func (g *Pin) Close() error {
	return nil
}
