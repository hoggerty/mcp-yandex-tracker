package issue

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/issue"
)

type SuggestTool struct{ Svc *logicissue.Service }

func (t *SuggestTool) Description() mcp.Tool {
	return mcp.NewTool("issue_suggest",
		mcp.WithDescription("Get issue suggestions matching input text, as used by Tracker's search-as-you-type."),
		mcp.WithString("input", mcp.Required(), mcp.Description("Text to match against issue keys and summaries")),
		mcp.WithString("queue", mcp.Description("Restrict suggestions to this queue key")),
	)
}

func (t *SuggestTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	input, err := req.RequireString("input")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return result.JSON(t.Svc.Suggest(ctx, input, req.GetString("queue", "")))
}
