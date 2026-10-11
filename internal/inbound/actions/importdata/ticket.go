package importdata

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/result"
	logicimportdata "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/importdata"
)

type TicketTool struct{ Svc *logicimportdata.Service }

func (t *TicketTool) Description() mcp.Tool {
	return mcp.NewTool("import_ticket",
		mcp.WithDescription("Create an issue preserving historical metadata (createdAt, createdBy, key). Administrator only."),
		mcp.WithObject("fields", mcp.Required(), mcp.Description("e.g. {\"queue\": \"TEST\", \"summary\": \"...\", \"createdAt\": \"2017-08-29T12:34:41.740+0000\", \"createdBy\": 11234567, \"key\": \"TEST-1\", \"description\": \"...\", \"assignee\": 11234567, \"priority\": 3, \"status\": 17}; queue, summary, createdAt, createdBy are required")),
	)
}

func (t *TicketTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fields, ok := req.GetArguments()["fields"].(map[string]any)
	if !ok || len(fields) == 0 {
		return mcp.NewToolResultError("fields must be a non-empty object"), nil
	}
	return result.JSON(t.Svc.Ticket(ctx, fields))
}
