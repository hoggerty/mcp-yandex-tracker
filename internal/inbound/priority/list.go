package priority

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicpriority "github.com/hoggerty/yandex-tracker-mcp/internal/logic/priority"
)

type listTool struct{ svc *logicpriority.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("priority_list",
		mcp.WithDescription("List the issue priorities available in the Yandex Tracker organization."),
		mcp.WithBoolean("localized", mcp.Description("If false, return names in all supported languages; default true")),
	)
}

func (t *listTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx, req.GetBool("localized", true)))
}
