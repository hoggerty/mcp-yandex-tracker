package externallink

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicexternallink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/externallink"
)

type deleteTool struct{ svc *logicexternallink.Service }

func (t *deleteTool) Description() mcp.Tool {
	return mcp.NewTool("externallink_delete",
		mcp.WithDescription("Remove a Yandex Tracker issue's link to an external application object."),
		mcp.WithDestructiveHintAnnotation(true),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("link_id", mcp.Required(), mcp.Description("External link id")),
	)
}

func (t *deleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	linkID, err := req.RequireString("link_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Delete(ctx, issueKey, linkID))
}
