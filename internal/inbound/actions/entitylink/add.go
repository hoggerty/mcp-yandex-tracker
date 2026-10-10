package entitylink

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/result"
	logicentitylink "github.com/hoggerty/yandex-tracker-mcp/internal/logic/actions/entitylink"
)

type AddTool struct{ Svc *logicentitylink.Service }

func (t *AddTool) Description() mcp.Tool {
	return mcp.NewTool("entitylink_add",
		mcp.WithDescription("Add one or more links from a Yandex Tracker entity to other entities or issues."),
		mcp.WithString("entity_type", mcp.Required(), mcp.Description("Entity type: project, portfolio, or goal")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Entity id or shortId")),
		mcp.WithObject("link", mcp.Description("Single link object, e.g. {\"relationship\": \"works towards\", \"entity\": \"ENTITY_ID\"}. Use this or links, not both.")),
		mcp.WithArray("links", mcp.Description("Array of link objects, for adding multiple links at once. Use this or link, not both.")),
	)
}

func (t *AddTool) Handle(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	entityType, err := req.RequireString("entity_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	id, err := req.RequireString("id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	args := req.GetArguments()
	var body any
	if link, ok := args["link"].(map[string]any); ok && len(link) > 0 {
		body = link
	} else if links, ok := args["links"].([]any); ok && len(links) > 0 {
		body = links
	} else {
		return mcp.NewToolResultError("either link or links must be provided"), nil
	}
	return result.JSON(t.Svc.Add(ctx, entityType, id, body))
}
