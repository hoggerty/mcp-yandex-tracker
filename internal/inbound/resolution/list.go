package resolution

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicresolution "github.com/hoggerty/yandex-tracker-mcp/internal/logic/resolution"
)

type listTool struct{ svc *logicresolution.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("resolution_list",
		mcp.WithDescription("List the issue resolutions available in the Yandex Tracker organization."),
	)
}

func (t *listTool) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx))
}
