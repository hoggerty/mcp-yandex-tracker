package sprint

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicsprint "github.com/hoggerty/yandex-tracker-mcp/internal/logic/sprint"
)

type startTool struct{ svc *logicsprint.Service }

func (t *startTool) Description() mcp.Tool {
	return mcp.NewTool("sprint_start",
		mcp.WithDescription("Start a Yandex Tracker sprint (set its status to in progress)."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Sprint id")),
		mcp.WithNumber("version", mcp.Description("Sprint's current version (for optimistic locking); omit to skip the check")),
	)
}

func (t *startTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Start(ctx, id, req.GetInt("version", 0)))
}
