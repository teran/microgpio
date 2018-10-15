package fake

import (
	"github.com/teran/microgpio/drivers"
)

var _ drivers.Driver = &FakeDriver{}

// FakeDriver type
type FakeDriver struct {
	ExportFunc   func() error
	UnexportFunc func() error
	LowFunc      func() error
	HighFunc     func() error
	InputFunc    func() error
	OutputFunc   func() error
	CloseFunc    func() error
}

// Export emulates exporting pin to userspace
func (f *FakeDriver) Export() error {
	return f.ExportFunc()
}

// Unexport emulates exporting pin to userspace
func (f *FakeDriver) Unexport() error {
	return f.UnexportFunc()
}

// Input emulates setting pin to input mode
func (f *FakeDriver) Input() error {
	return f.InputFunc()
}

// Output emulates setting pin to output mode
func (f *FakeDriver) Output() error {
	return f.OutputFunc()
}

// Low emulates setting low bit to the pin specified
func (f *FakeDriver) Low() error {
	return f.LowFunc()
}

// High emulates setting high bit to the pin specified
func (f *FakeDriver) High() error {
	return f.HighFunc()
}

func (f *FakeDriver) Close() error {
	return f.CloseFunc()
}
