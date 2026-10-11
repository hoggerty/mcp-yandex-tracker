package version

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicversion "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/version"
)

type GetTool struct{ Svc *logicversion.Service }

func (t *GetTool) Description() mcp.Tool {
	return mcp.NewTool("version_get",
		mcp.WithDescription("Get the parameters of a single Yandex Tracker version."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Version id")),
	)
}

func (t *GetTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Get(ctx, id))
}
