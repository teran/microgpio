package models

// Status type
type Status string

const (
	// StatusOK means ok status
	StatusOK Status = "ok"

	// StatusError means error status
	StatusError Status = "error"
)

// EmptyStatus type
type EmptyStatus struct {
	Status Status `json:"status"`
}

// Mode type
type Mode string

const (
	// ModeIn means in mode for the pin
	ModeIn Mode = "in"

	// ModeOut means out mode for the pin
	ModeOut Mode = "out"
)

// StatusWithMode type
type StatusWithMode struct {
	Status Status `json:"status"`
	Mode   Mode   `json:"mode"`
}

// StatusWithValue type
type StatusWithValue struct {
	Status Status `json:"status"`
	Value  int    `json:"value"`
}
