package status

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicstatus "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/status"
)

type ListTool struct{ Svc *logicstatus.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("status_list",
		mcp.WithDescription("List the issue statuses available in the Yandex Tracker organization."),
	)
}

func (t *ListTool) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx))
}
