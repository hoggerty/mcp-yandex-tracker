package filter

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicfilter "github.com/hoggerty/yandex-tracker-mcp/internal/logic/filter"
)

type getTool struct{ svc *logicfilter.Service }

func (t *getTool) Description() mcp.Tool {
	return mcp.NewTool("filter_get",
		mcp.WithDescription("Get a saved issue filter by id."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Filter id")),
	)
}

func (t *getTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Get(ctx, id))
}
