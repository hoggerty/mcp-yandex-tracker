// Package result formats Tracker API responses as MCP tool results.
package result

import (
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// JSON renders a Tracker API call's result as a pretty-printed MCP tool result.
func JSON(raw json.RawMessage, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultErrorFromErr("tracker API call failed", err), nil
	}
	pretty, err := formatJSON(raw)
	if err != nil {
		return mcp.NewToolResultErrorFromErr("failed to format tracker API response", err), nil
	}
	return mcp.NewToolResultText(pretty), nil
}

func formatJSON(raw json.RawMessage) (string, error) {
	buf, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", fmt.Errorf("format JSON: %w", err)
	}
	return string(buf), nil
}
