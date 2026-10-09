package externallink

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicexternallink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/externallink"
)

type addTool struct{ svc *logicexternallink.Service }

func (t *addTool) Description() mcp.Tool {
	return mcp.NewTool("externallink_add",
		mcp.WithDescription("Link a Yandex Tracker issue to an object in an external application."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("relationship", mcp.Required(), mcp.Description("Relationship type, e.g. RELATES")),
		mcp.WithString("key", mcp.Required(), mcp.Description("External object key")),
		mcp.WithString("origin", mcp.Required(), mcp.Description("External application id, see externallink_applications_list")),
	)
}

func (t *addTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	relationship, err := req.RequireString("relationship")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	key, err := req.RequireString("key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	origin, err := req.RequireString("origin")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Add(ctx, issueKey, relationship, key, origin))
}
