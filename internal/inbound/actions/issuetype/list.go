package issuetype

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissuetype "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issuetype"
)

type ListTool struct{ Svc *logicissuetype.Service }

func (t *ListTool) Description() mcp.Tool {
	return mcp.NewTool("issuetype_list",
		mcp.WithDescription("List the issue types available in the Yandex Tracker organization."),
	)
}

func (t *ListTool) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.List(ctx))
}
