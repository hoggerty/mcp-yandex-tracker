// Package safety classifies what a Tracker operation does to data. Each
// logic operation declares its own Level next to its implementation, so the
// classification is a property of the business operation itself rather than
// of how it happens to be exposed as an MCP tool.
package safety

import "errors"

type Level int

const (
	// ReadOnly operations do not modify Tracker data.
	ReadOnly Level = iota
	// Write operations modify Tracker data but are not destructive
	// (create, edit, add, move, transition, restore, ...).
	Write
	// Destructive operations delete data or act on many items at once
	// (delete, bulk change, ...).
	Destructive
)

// Mode is how much of the Tracker API the server exposes. It is the single
// source of truth for the mode/level relationship: both tool registration
// (inbound) and operation execution (logic) consult Mode.Allows, so a
// disabled operation is unavailable at every layer, not just hidden from the
// MCP tool list.
type Mode string

const (
	ModeReadOnly  Mode = "read-only"
	ModeReadWrite Mode = "read-write"
	ModeFull      Mode = "full"
)

// Valid reports whether m is one of the known modes.
func (m Mode) Valid() bool {
	switch m {
	case ModeReadOnly, ModeReadWrite, ModeFull:
		return true
	default:
		return false
	}
}

// Allows reports whether an operation at the given Level may run under m.
func (m Mode) Allows(level Level) bool {
	switch m {
	case ModeReadOnly:
		return level == ReadOnly
	case ModeFull:
		return true
	default: // ModeReadWrite
		return level != Destructive
	}
}

// ErrDisabled is returned by a logic operation whose Level is not permitted
// by the server's current Mode.
var ErrDisabled = errors.New("operation disabled by tools mode")

// Guard is held by a logic Service and checked by every operation before it
// touches the Tracker API, so a mode-disabled operation is unavailable at
// the logic layer itself — not merely absent from the registered MCP tools.
type Guard struct{ mode Mode }

// NewGuard builds a Guard enforcing mode.
func NewGuard(mode Mode) Guard { return Guard{mode: mode} }

// Check returns ErrDisabled if level is not permitted by the guard's mode.
func (g Guard) Check(level Level) error {
	if !g.mode.Allows(level) {
		return ErrDisabled
	}
	return nil
}
