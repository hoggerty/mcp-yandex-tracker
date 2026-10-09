package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// GetSprint retrieves the parameters of a single sprint.
func (c *Client) GetSprint(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/sprints/"+url.PathEscape(id), nil, nil)
}

// CreateSprint creates a new sprint on a board.
func (c *Client) CreateSprint(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/sprints", nil, fields)
}

// EditSprint edits a sprint; version must match the sprint's current version.
func (c *Client) EditSprint(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/sprints/"+url.PathEscape(id), q, fields)
}

// StartSprint changes a sprint's status to "in progress". version is optional (0 to omit).
func (c *Client) StartSprint(ctx context.Context, id string, version int) (json.RawMessage, error) {
	q := url.Values{}
	if version > 0 {
		q.Set("version", strconv.Itoa(version))
	}
	return c.do(ctx, http.MethodPost, "/v3/sprints/"+url.PathEscape(id)+"/_start", q, nil)
}

// ArchiveSprint archives a sprint. version is optional (0 to omit).
func (c *Client) ArchiveSprint(ctx context.Context, id string, version int) (json.RawMessage, error) {
	q := url.Values{}
	if version > 0 {
		q.Set("version", strconv.Itoa(version))
	}
	return c.do(ctx, http.MethodPost, "/v3/sprints/"+url.PathEscape(id)+"/_archive", q, nil)
}

// DeleteSprint deletes a sprint.
func (c *Client) DeleteSprint(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/sprints/"+url.PathEscape(id), nil, nil)
}
