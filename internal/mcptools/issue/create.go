package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/mcptools/result"
	"github.com/hoggerty/yandex-tracker-mcp/internal/trackerapi"
)

type createTool struct{ client *trackerapi.Client }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("issue_create",
		mcp.WithDescription("Create a new Yandex Tracker issue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key the issue should be created in, e.g. QUEUE")),
		mcp.WithString("summary", mcp.Required(), mcp.Description("Issue title")),
		mcp.WithString("description", mcp.Description("Issue description")),
		mcp.WithObject("fields", mcp.Description("Additional raw fields to set, merged over queue/summary/description")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queue, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	summary, err := req.RequireString("summary")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	fields := map[string]any{
		"queue":   map[string]any{"key": queue},
		"summary": summary,
	}
	if desc := req.GetString("description", ""); desc != "" {
		fields["description"] = desc
	}
	if extra, ok := req.GetArguments()["fields"].(map[string]any); ok {
		for k, v := range extra {
			fields[k] = v
		}
	}
	return result.JSON(t.client.CreateIssue(ctx, fields))
}
