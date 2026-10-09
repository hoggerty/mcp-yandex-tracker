package externallink

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicexternallink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/externallink"
)

type applicationsListTool struct{ svc *logicexternallink.Service }

func (t *applicationsListTool) Description() mcp.Tool {
	return mcp.NewTool("externallink_applications_list",
		mcp.WithDescription("List external applications known to Tracker (usable as external-link origins)."),
		mcp.WithReadOnlyHintAnnotation(true),
	)
}

func (t *applicationsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.svc.ApplicationsList(ctx))
}
