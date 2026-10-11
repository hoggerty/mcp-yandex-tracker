package queue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicqueue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/queue"
)

type TagDeleteTool struct{ Svc *logicqueue.Service }

func (t *TagDeleteTool) Description() mcp.Tool {
	return mcp.NewTool("queue_tag_delete",
		mcp.WithDescription("Remove a tag from a Yandex Tracker queue. Administrator-only; the tag must be unused by any issue."),
		mcp.WithString("queue", mcp.Required(), mcp.Description("Queue key or id")),
		mcp.WithString("tag", mcp.Required(), mcp.Description("Tag name")),
	)
}

func (t *TagDeleteTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	queueKey, err := req.RequireString("queue")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	tag, err := req.RequireString("tag")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.TagDelete(ctx, queueKey, tag))
}
