// Package tool defines the interface Yandex Tracker MCP tools implement.
package tool

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Tool is a single Yandex Tracker operation exposed as an MCP tool. Each
// implementation is a standalone struct so that a tool's definition, schema,
// and handler live together.
type Tool interface {
	Description() mcp.Tool
	Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// NewServerTool adapts a Tool to the type the mcp-go server registers.
func NewServerTool(t Tool) server.ServerTool {
	return server.ServerTool{Tool: t.Description(), Handler: t.Handle}
}
