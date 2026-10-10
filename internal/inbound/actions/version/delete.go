package version

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/version"
)

type DeleteTool struct{ Svc *logicversion.Service }

func (t *DeleteTool) Description() mcp.Tool {
	return mcp.NewTool("version_delete",
		mcp.WithDescription("Delete a Yandex Tracker version."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Version id")),
	)
}

func (t *DeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Delete(ctx, id))
}
