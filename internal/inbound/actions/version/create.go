package version

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/version"
)

type CreateTool struct{ Svc *logicversion.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("version_create",
		mcp.WithDescription("Create a new version on a Yandex Tracker queue."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Version definition, e.g. {\"queue\": \"TEST\", \"name\": \"1.0\", \"description\": \"...\", \"startDate\": \"2026-01-01\", \"dueDate\": \"2026-06-01\"}")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Create(ctx, fields))
}
