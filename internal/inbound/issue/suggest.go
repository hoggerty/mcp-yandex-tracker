package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

type suggestTool struct{ svc *logicissue.Service }

func (t *suggestTool) Description() mcp.Tool {
	return mcp.NewTool("issue_suggest",
		mcp.WithDescription("Get issue suggestions matching input text, as used by Tracker's search-as-you-type."),
		mcp.WithReadOnlyHintAnnotation(true),
		mcp.WithString("input", mcp.Required(), mcp.Description("Text to match against issue keys and summaries")),
		mcp.WithString("queue", mcp.Description("Restrict suggestions to this queue key")),
	)
}

func (t *suggestTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	input, err := req.RequireString("input")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.svc.Suggest(ctx, input, req.GetString("queue", "")))
}
