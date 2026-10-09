package sprint

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicsprint "github.com/hoggerty/yandex-tracker-mcp/internal/logic/sprint"
)

type createTool struct{ svc *logicsprint.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("sprint_create",
		mcp.WithDescription("Create a new sprint on a Yandex Tracker agile board."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Sprint definition, e.g. {\"name\": \"...\", \"board\": {\"id\": \"1\"}, \"startDate\": \"2024-01-01\", \"endDate\": \"2024-01-14\"}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
