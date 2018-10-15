package drivers

// Driver type
type Driver interface {
	High() error
	Low() error
	Input() error
	Output() error
	Close() error
}
