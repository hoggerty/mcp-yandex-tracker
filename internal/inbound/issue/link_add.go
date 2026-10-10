package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type linkAddTool struct{ svc *logicissue.Service }

func (t *linkAddTool) Description() mcp.Tool {
	return mcp.NewTool("issue_link_add",
		mcp.WithDescription("Create a link between two Yandex Tracker issues."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Source issue key")),
		mcp.WithString("relationship", mcp.Required(), mcp.Description("Relationship type, e.g. \"relates\", \"is dependent by\", \"depends on\"")),
		mcp.WithString("target_issue_key", mcp.Required(), mcp.Description("Issue key to link to")),
	)
}

func (t *linkAddTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	relationship, err := req.RequireString("relationship")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	target, err := req.RequireString("target_issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.LinkIssues(ctx, issueKey, relationship, target))
}
