package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issue"
)

type CreateTool struct{ Svc *logicissue.Service }

func (t *CreateTool) Description() mcp.Tool {
	return mcp.NewTool("issue_create",
		mcp.WithDescription("Create a new Yandex Tracker issue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key the issue should be created in, e.g. QUEUE")),
		mcp.WithString("summary", mcp.Required(), mcp.Description("Issue title")),
		mcp.WithString("description", mcp.Description("Issue description")),
		mcp.WithObject("fields", mcp.Description("Additional raw fields to set, merged over queue/summary/description")),
	)
}

func (t *CreateTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queue, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	summary, err := req.RequireString("summary")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	extra, _ := req.GetArguments()["fields"].(map[string]any)
	return result.JSON(t.Svc.Create(ctx, queue, summary, req.GetString("description", ""), extra))
}
