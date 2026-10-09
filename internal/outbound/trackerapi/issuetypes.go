package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListIssueTypes lists the issue types available in the organization.
func (c *Client) ListIssueTypes(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issuetypes", nil, nil)
}

// CreateIssueType creates a new issue type.
func (c *Client) CreateIssueType(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issuetypes", nil, fields)
}

// EditIssueType edits an issue type; version must match the issue type's current version.
func (c *Client) EditIssueType(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/issuetypes/"+url.PathEscape(id), q, fields)
}
