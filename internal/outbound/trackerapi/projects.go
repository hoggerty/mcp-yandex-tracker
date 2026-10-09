package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// CreateProject creates a new project.
func (c *Client) CreateProject(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/projects", nil, fields)
}

// GetProject retrieves a single project by id.
func (c *Client) GetProject(ctx context.Context, id, expand string) (json.RawMessage, error) {
	q := url.Values{}
	if expand != "" {
		q.Set("expand", expand)
	}
	return c.do(ctx, http.MethodGet, "/v3/projects/"+url.PathEscape(id), q, nil)
}

// ListProjects lists all projects.
func (c *Client) ListProjects(ctx context.Context, expand string) (json.RawMessage, error) {
	q := url.Values{}
	if expand != "" {
		q.Set("expand", expand)
	}
	return c.do(ctx, http.MethodGet, "/v3/projects", q, nil)
}

// ListProjectQueues lists the queues attached to a project.
func (c *Client) ListProjectQueues(ctx context.Context, id, expand string) (json.RawMessage, error) {
	q := url.Values{}
	if expand != "" {
		q.Set("expand", expand)
	}
	return c.do(ctx, http.MethodGet, "/v3/projects/"+url.PathEscape(id)+"/queues", q, nil)
}

// UpdateProject replaces a project's fields.
func (c *Client) UpdateProject(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPut, "/v3/projects/"+url.PathEscape(id), q, fields)
}

// DeleteProject deletes a project.
func (c *Client) DeleteProject(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/projects/"+url.PathEscape(id), nil, nil)
}
