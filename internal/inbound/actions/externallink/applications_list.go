package externallink

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicexternallink "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/externallink"
)

type ApplicationsListTool struct{ Svc *logicexternallink.Service }

func (t *ApplicationsListTool) Description() mcp.Tool {
	return mcp.NewTool("externallink_applications_list",
		mcp.WithDescription("List external applications known to Tracker (usable as external-link origins)."),
	)
}

func (t *ApplicationsListTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.ApplicationsList(ctx))
}
