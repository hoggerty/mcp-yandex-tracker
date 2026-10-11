package externallink

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicexternallink "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/externallink"
)

type DeleteTool struct{ Svc *logicexternallink.Service }

func (t *DeleteTool) Description() mcp.Tool {
	return mcp.NewTool("externallink_delete",
		mcp.WithDescription("Remove a Yandex Tracker issue's link to an external application object."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("link_id", mcp.Required(), mcp.Description("External link id")),
	)
}

func (t *DeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	linkID, err := req.RequireString("link_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Delete(ctx, issueKey, linkID))
}
