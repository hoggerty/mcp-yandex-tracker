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
// and handler live together.
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

// ServerTool adapts Entry to the type the mcp-go server registers.
func (e Entry) ServerTool() server.ServerTool {
	return server.ServerTool{Tool: e.Tool.Description(), Handler: e.Tool.Handle}
}
