// Package tool defines the interface Yandex Tracker MCP tools implement.
package tool

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

// Tool is a single Yandex Tracker operation exposed as an MCP tool. Each
// implementation is a standalone struct so that a tool's definition, schema,
// and handler live together. Description must not set the ReadOnlyHint or
// DestructiveHint annotations itself — Entry derives them from the logic
// layer's Safety classification, so they can't drift out of sync with it.
type Tool interface {
	Description() mcp.Tool
	Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Entry pairs a Tool with the safety classification of the logic operation
// it wraps. The classification is declared in the logic layer, next to that
// operation's implementation — a domain's Tools() function only relays it.
type Entry struct {
	Tool   Tool
	Safety safety.Level
}

// ServerTool adapts Entry to the type the mcp-go server registers, setting
// the tool's ReadOnlyHint/DestructiveHint annotations from Safety.
func (e Entry) ServerTool() server.ServerTool {
	t := e.Tool.Description()
	readOnly := e.Safety == safety.ReadOnly
	destructive := e.Safety == safety.Destructive
	t.Annotations.ReadOnlyHint = &readOnly
	t.Annotations.DestructiveHint = &destructive
	return server.ServerTool{Tool: t, Handler: e.Tool.Handle}
}
