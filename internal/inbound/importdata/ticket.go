package importdata

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicimportdata "github.com/hoggerty/yandex-tracker-mcp/internal/logic/importdata"
)

type ticketTool struct{ svc *logicimportdata.Service }

func (t *ticketTool) Description() mcp.Tool {
	return mcp.NewTool("import_ticket",
		mcp.WithDescription("Create an issue preserving historical metadata (createdAt, createdBy, key). Administrator only."),
		mcp.WithDestructiveHintAnnotation(false),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"queue\": \"TEST\", \"summary\": \"...\", \"createdAt\": \"2017-08-29T12:34:41.740+0000\", \"createdBy\": 11234567, \"key\": \"TEST-1\", \"description\": \"...\", \"assignee\": 11234567, \"priority\": 3, \"status\": 17}; queue, summary, createdAt, createdBy are required")),
	)
}

func (t *ticketTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.svc.Ticket(ctx, fields))
}
