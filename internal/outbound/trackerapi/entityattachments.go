package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// ListEntityAttachments lists the files attached to an entity.
func (c *Client) ListEntityAttachments(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/attachments"
	return c.do(ctx, http.MethodGet, path, nil, nil)
}

// GetEntityAttachment retrieves metadata of a single file attached to an entity.
func (c *Client) GetEntityAttachment(ctx context.Context, entityType, id, fileID string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/attachments/" + url.PathEscape(fileID)
	return c.do(ctx, http.MethodGet, path, nil, nil)
}

// DeleteEntityAttachment deletes a file attached to an entity.
func (c *Client) DeleteEntityAttachment(ctx context.Context, entityType, id, fileID string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/attachments/" + url.PathEscape(fileID)
	return c.do(ctx, http.MethodDelete, path, nil, nil)
}
