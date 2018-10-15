package drivers

// Driver type
type Driver interface {
	Export() error
	Unexport() error
	Input() error
	Output() error
	Low() error
	High() error
	Close() error
}
