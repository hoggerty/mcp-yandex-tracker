package issuetype

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissuetype "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issuetype"
)

type listTool struct{ svc *logicissuetype.Service }

func (t *listTool) Description() mcp.Tool {
	return mcp.NewTool("issuetype_list",
		mcp.WithDescription("List the issue types available in the Yandex Tracker organization."),
		mcp.WithReadOnlyHintAnnotation(true),
	)
}

func (t *listTool) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.List(ctx))
}
