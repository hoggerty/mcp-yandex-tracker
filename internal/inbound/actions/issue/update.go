package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issue"
)

type UpdateTool struct{ Svc *logicissue.Service }

func (t *UpdateTool) Description() mcp.Tool {
	return mcp.NewTool("issue_update",
		mcp.WithDescription("Update fields on an existing Yandex Tracker issue."),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Fields to update, e.g. {\"summary\": \"new title\", \"priority\": {\"id\": \"2\"}}")),
	)
}

func (t *UpdateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Update(ctx, issueKey, fields))
}
