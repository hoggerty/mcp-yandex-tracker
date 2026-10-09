package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type checklistAddTool struct{ svc *logicissue.Service }

func (t *checklistAddTool) Description() mcp.Tool {
	return mcp.NewTool("issue_checklist_add",
		mcp.WithDescription("Add an item to a Yandex Tracker issue's checklist, creating the checklist if needed."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithString("issue_key", mcp.Required(), mcp.Description("Issue key, e.g. QUEUE-123")),
		mcp.WithString("text", mcp.Required(), mcp.Description("Checklist item text")),
		mcp.WithObject("fields", mcp.Description("Additional item fields, e.g. assignee, deadline")),
	)
}

func (t *checklistAddTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	issueKey, err := req.RequireString("issue_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	text, err := req.RequireString("text")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	item, _ := req.GetArguments()["fields"].(map[string]any)
	if item == nil {
		item = map[string]any{}
	}
	item["text"] = text
	return result.JSON(t.svc.ChecklistAdd(ctx, issueKey, item))
}
