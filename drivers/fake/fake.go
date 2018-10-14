package fake

import (
	"github.com/teran/microgpio/drivers"
)

var _ drivers.Driver = &FakeDriver{}

// FakeDriver type
type FakeDriver struct {
	LowFunc    func(int) error
	HighFunc   func(int) error
	OutputFunc func(int) error
}

// Output emulates setting pin to output mode
func (f *FakeDriver) Output(id int) error {
	return f.OutputFunc(id)
}

// Low emulates setting low bit to the pin specified
func (f *FakeDriver) Low(id int) error {
	return f.LowFunc(id)
}

// High emulates setting high bit to the pin specified
func (f *FakeDriver) High(id int) error {
	return f.HighFunc(id)
}
