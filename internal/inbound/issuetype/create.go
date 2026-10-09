package issuetype

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicissuetype "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issuetype"
)

type createTool struct{ svc *logicissuetype.Service }

func (t *createTool) Description() mcp.Tool {
	return mcp.NewTool("issuetype_create",
		mcp.WithDescription("Create a new issue type (requires Administrator rights)."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("Issue type definition, e.g. {\"key\": \"client\", \"name\": {\"ru\": \"Клиент\", \"en\": \"Customer\"}}")),
	)
}

func (t *createTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Create(ctx, fields))
}
