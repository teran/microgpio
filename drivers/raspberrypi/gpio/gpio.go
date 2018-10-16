package gpio

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"sync"

	"github.com/teran/microgpio/drivers"
	"github.com/teran/microgpio/models"
)

var _ drivers.Driver = (*Pin)(nil)

// Files to use for operations(relative pathes)
const (
	ExportFilepath    = "class/gpio/export"
	UnexportFilepath  = "class/gpio/unexport"
	DirectionFilepath = "class/gpio/gpio%d/direction"
	ValueFilepath     = "class/gpio/gpio%d/value"
)

// Pin type
type Pin struct {
	sync.Mutex

	id                int
	exportFilepath    string
	unexportFilepath  string
	directionFilepath string
	valueFilepath     string
}

// New returns new *GPIOPin instance
func New(id int) *Pin {
	return NewWithPrefix(id, "/sys")
}

// NewWithPrefix allows to specify prefix to use for sysfs
func NewWithPrefix(id int, prefix string) *Pin {
	return &Pin{
		id:                id,
		exportFilepath:    path.Join(prefix, ExportFilepath),
		unexportFilepath:  path.Join(prefix, UnexportFilepath),
		directionFilepath: path.Join(prefix, DirectionFilepath),
		valueFilepath:     path.Join(prefix, ValueFilepath),
	}
}

// Export exports pin to userspace
func (g *Pin) Export() error {
	return ioutil.WriteFile(
		g.exportFilepath,
		[]byte(strconv.Itoa(g.id)),
		0644,
	)
}

// Unexport unexports pin from userspace
func (g *Pin) Unexport() error {
	return ioutil.WriteFile(
		g.unexportFilepath,
		[]byte(strconv.Itoa(g.id)),
		0644,
	)
}

// Input sets input mode for the pin
func (g *Pin) Input() error {
	return ioutil.WriteFile(
		fmt.Sprintf(g.directionFilepath, g.id),
		[]byte("in"),
		0644,
	)
}

// Output sets output mode for the pin
func (g *Pin) Output() error {
	return ioutil.WriteFile(
		fmt.Sprintf(g.directionFilepath, g.id),
		[]byte("out"),
		0644,
	)
}

// Mode returns pin mode. Normally values should be on of: in, out
func (g *Pin) Mode() (models.Mode, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf(g.directionFilepath, g.id))
	if err != nil {
		return "", err
	}
	return models.Mode(data), err
}

// High sets high bit for GPIO pin
func (g *Pin) High() error {
	return ioutil.WriteFile(
		fmt.Sprintf(g.valueFilepath, g.id),
		[]byte("1"),
		0644,
	)
}

// Low sets low bit to GPIO pin
func (g *Pin) Low() error {
	return ioutil.WriteFile(
		fmt.Sprintf(g.valueFilepath, g.id),
		[]byte("0"),
		0644,
	)
}

// Value returns current value set for the pin
func (g *Pin) Value() (int, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf(g.valueFilepath, g.id))
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(data))
}

// Close should destroy all the resources allocated by the *Pin object
func (g *Pin) Close() error {
	return nil
}
