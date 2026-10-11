package sprint

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicsprint "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/sprint"
)

type DeleteTool struct{ Svc *logicsprint.Service }

func (t *DeleteTool) Description() mcp.Tool {
	return mcp.NewTool("sprint_delete",
		mcp.WithDescription("Delete a Yandex Tracker sprint."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Sprint id")),
	)
}

func (t *DeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Delete(ctx, id))
}
