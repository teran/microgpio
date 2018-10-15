package drivers

// Driver type
type Driver interface {
	Close() error
	Export() error
	High() error
	Input() error
	Low() error
	Mode() (string, error)
	Output() error
	Unexport() error
	Value() (int, error)
}
