package models

// Mode type
type Mode string

const (
	// ModeIn means in mode for the pin
	ModeIn Mode = "in"

	// ModeOut means out mode for the pin
	ModeOut Mode = "out"
)

// ToggleStatus type
type ToggleStatus string

const (
	// ToggleStatusOn displays the pin is On
	ToggleStatusOn ToggleStatus = "on"

	// ToggleStatusOff displays the pin is Off
	ToggleStatusOff ToggleStatus = "off"

	// ToggleStatusUnknown display the pin status is Unknown
	ToggleStatusUnknown ToggleStatus = "unknown"
)

// Status model
type Status struct {
	Status ToggleStatus `json:"status"`
}

// PingStatus type
type PingStatus string

const (
	// PingStatusOK returned on ok status for ping
	PingStatusOK PingStatus = "ok"
)

// PingResponse model
type PingResponse struct {
	Status PingStatus `json:"status"`
}
