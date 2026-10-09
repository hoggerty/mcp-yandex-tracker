package version

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/version"
)

type editTool struct{ svc *logicversion.Service }

func (t *editTool) Description() mcp.Tool {
	return mcp.NewTool("version_edit",
		mcp.WithDescription("Edit a Yandex Tracker version."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Version id")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": \"...\", \"description\": \"...\", \"startDate\": \"2026-01-01\", \"dueDate\": \"2026-06-01\"}")),
	)
}

func (t *editTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Edit(ctx, id, fields))
}
