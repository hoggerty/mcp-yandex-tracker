package version

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/version"
)

type deleteTool struct{ svc *logicversion.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("version_delete",
		mcp.WithDescription("Delete a Yandex Tracker version."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("id", mcp.Required(), mcp.Description("Version id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, id))
}
