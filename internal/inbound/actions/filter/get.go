package filter

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicfilter "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/filter"
)

type GetTool struct{ Svc *logicfilter.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("filter_get",
		mcp.WithDescription("Get a saved issue filter by id."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Filter id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, id))
}
