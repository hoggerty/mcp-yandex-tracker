package sprint

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicsprint "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/sprint"
)

type GetTool struct{ Svc *logicsprint.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("sprint_get",
		mcp.WithDescription("Get a Yandex Tracker sprint by id."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Sprint id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, id))
}
