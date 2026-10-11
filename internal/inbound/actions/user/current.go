package user

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicuser "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/user"
)

type CurrentTool struct{ Svc *logicuser.Service }

func (t *CurrentTool) Description() mcp.Tool {
	return mcp.NewTool("user_current",
		mcp.WithDescription("Retrieve the Yandex Tracker profile of the user that owns the configured API token."),
	)
}

func (t *CurrentTool) Handle(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return result.JSON(t.Svc.Current(ctx))
}
