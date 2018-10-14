package drivers

// Driver type
type Driver interface {
	High(int) error
	Low(int) error
	Output(int) error
}
