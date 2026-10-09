package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// AddEntityChecklistItem adds an item to an entity's checklist (creating the checklist if needed).
func (c *Client) AddEntityChecklistItem(ctx context.Context, entityType, id string, fields map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/checklistItems"
	return c.do(ctx, http.MethodPost, path, nil, fields)
}

// EditEntityChecklist bulk-edits an entity's checklist items.
func (c *Client) EditEntityChecklist(ctx context.Context, entityType, id string, items []map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/checklistItems"
	return c.do(ctx, http.MethodPatch, path, nil, items)
}

// EditEntityChecklistItem edits a single checklist item of an entity.
func (c *Client) EditEntityChecklistItem(ctx context.Context, entityType, id, itemID string, fields map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/checklistItems/" + url.PathEscape(itemID)
	return c.do(ctx, http.MethodPatch, path, nil, fields)
}

// MoveEntityChecklistItem moves a checklist item so it is positioned before the given item.
func (c *Client) MoveEntityChecklistItem(ctx context.Context, entityType, id, itemID, beforeID string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/checklistItems/" + url.PathEscape(itemID) + "/_move"
	return c.do(ctx, http.MethodPost, path, nil, map[string]any{"before": beforeID})
}

// DeleteEntityChecklist deletes all checklist items of an entity.
func (c *Client) DeleteEntityChecklist(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/checklistItems"
	return c.do(ctx, http.MethodDelete, path, nil, nil)
}

// DeleteEntityChecklistItem deletes a single checklist item of an entity.
func (c *Client) DeleteEntityChecklistItem(ctx context.Context, entityType, id, itemID string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/checklistItems/" + url.PathEscape(itemID)
	return c.do(ctx, http.MethodDelete, path, nil, nil)
}
