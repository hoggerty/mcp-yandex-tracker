package status

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicstatus "github.com/hoggerty/yandex-tracker-mcp/internal/logic/status"
)

type listTool struct{ svc *logicstatus.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("status_list",
		mcp.WithDescription("List the issue statuses available in the Yandex Tracker organization."),
		mcp.WithReadOnlyHintAnnotation(true),
	)
}

func (t *listTool) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx))
}
