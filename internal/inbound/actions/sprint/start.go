package sprint

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicsprint "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/sprint"
)

type StartTool struct{ Svc *logicsprint.Service }

func (t *StartTool) Description() mcp.Tool {
	return mcp.NewTool("sprint_start",
		mcp.WithDescription("Start a Yandex Tracker sprint (set its status to in progress)."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Sprint id")),
		mcp.WithNumber("version", mcp.Description("Sprint's current version (for optimistic locking); omit to skip the check")),
	)
}

func (t *StartTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Start(ctx, id, req.GetInt("version", 0)))
}
