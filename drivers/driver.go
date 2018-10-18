package drivers

import (
	"github.com/teran/microgpio/models"
)

// Driver type
type Driver interface {
	Close() error
	Export() error
	IsExported() bool
	High() error
	Input() error
	Low() error
	Mode() (models.Mode, error)
	Output() error
	Unexport() error
	Value() (int, error)
}
