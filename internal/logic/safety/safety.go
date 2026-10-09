// Package safety classifies what a Tracker operation does to data. Each
// logic operation declares its own Level next to its implementation, so the
// classification is a property of the business operation itself rather than
// of how it happens to be exposed as an MCP tool.
package safety

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
