package version

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicversion "github.com/hoggerty/yandex-tracker-mcp/internal/logic/version"
)

type getTool struct{ svc *logicversion.Service }

func (t *getTool) Description() mcp.Tool {
	return mcp.NewTool("version_get",
		mcp.WithDescription("Get the parameters of a single Yandex Tracker version."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("id", mcp.Required(), mcp.Description("Version id")),
	)
}

func (t *getTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Get(ctx, id))
}
