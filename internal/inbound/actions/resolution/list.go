package resolution

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicresolution "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/resolution"
)

type ListTool struct{ Svc *logicresolution.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("resolution_list",
		mcp.WithDescription("List the issue resolutions available in the Yandex Tracker organization."),
	)
}

func (t *ListTool) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx))
}
