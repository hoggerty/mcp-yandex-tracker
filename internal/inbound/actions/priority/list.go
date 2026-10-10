package priority

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicpriority "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/priority"
)

type ListTool struct{ Svc *logicpriority.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("priority_list",
		mcp.WithDescription("List the issue priorities available in the Yandex Tracker organization."),
		mcp.WithBoolean("localized", mcp.Description("If false, return names in all supported languages; default true")),
	)
}

func (t *ListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx, req.GetBool("localized", true)))
}
