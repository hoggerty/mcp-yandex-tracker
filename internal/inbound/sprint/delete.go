package sprint

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicsprint "github.com/hoggerty/yandex-tracker-mcp/internal/logic/sprint"
)

type deleteTool struct{ svc *logicsprint.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("sprint_delete",
		mcp.WithDescription("Delete a Yandex Tracker sprint."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Sprint id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, id))
}
