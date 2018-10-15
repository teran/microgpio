package drivers

// Driver type
type Driver interface {
	Export() error
	Unexport() error
	Input() error
	Output() error
	Mode() (string, error)
	Value() (int, error)
	Low() error
	High() error
	Close() error
}
