package status

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicstatus "github.com/hoggerty/yandex-tracker-mcp/internal/logic/status"
)

type createTool struct{ svc *logicstatus.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("status_create",
		mcp.WithDescription("Create a new issue status (requires Administrator rights)."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Status definition, e.g. {\"key\": \"myStatus\", \"name\": {\"ru\": \"...\", \"en\": \"...\"}, \"type\": \"paused\"}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
