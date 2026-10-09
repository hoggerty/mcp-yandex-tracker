package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// AddEntityLinks links an entity to one or more other entities. body may be a
// single {"relationship": ..., "entity": ...} object or an array of such objects.
func (c *Client) AddEntityLinks(ctx context.Context, entityType, id string, body any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/links"
	return c.do(ctx, http.MethodPost, path, nil, body)
}

// ListEntityLinks lists the entities linked to an entity. fields is optional (comma-separated linked entity fields).
func (c *Client) ListEntityLinks(ctx context.Context, entityType, id, fields string) (json.RawMessage, error) {
	q := url.Values{}
	if fields != "" {
		q.Set("fields", fields)
	}
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/links"
	return c.do(ctx, http.MethodGet, path, q, nil)
}

// DeleteEntityLink removes the link between an entity and the entity identified by rightID.
func (c *Client) DeleteEntityLink(ctx context.Context, entityType, id, rightID string) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("right", rightID)
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/links"
	return c.do(ctx, http.MethodDelete, path, q, nil)
}
