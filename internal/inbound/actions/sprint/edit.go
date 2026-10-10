package sprint

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicsprint "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/sprint"
)

type EditTool struct{ Svc *logicsprint.Service }

func (t *EditTool) Description() mcp.Tool {
	return mcp.NewTool("sprint_edit",
		mcp.WithDescription("Edit a Yandex Tracker sprint."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Sprint id")),
		mcp.WithNumber("version", mcp.Required(), mcp.Description("Sprint's current version (for optimistic locking)")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Parameters to set, e.g. {\"name\": \"...\", \"startDate\": \"2024-01-01\", \"endDate\": \"2024-01-14\"}")),
	)
}

func (t *EditTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	version, err := req.RequireInt("version")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Edit(ctx, id, version, fields))
}
