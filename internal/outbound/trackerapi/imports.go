package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// ImportTicket creates an issue preserving historical metadata (createdAt, createdBy, key). Administrator only.
func (c *Client) ImportTicket(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issues/_import", nil, fields)
}

// ImportComment adds a comment to an issue preserving historical metadata. Administrator only.
func (c *Client) ImportComment(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueID)+"/comments/_import", nil, fields)
}

// ImportLink adds a link to an issue preserving historical metadata. Administrator only.
func (c *Client) ImportLink(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueID)+"/links/_import", nil, fields)
}

// ImportWorklog adds a worklog entry to an issue preserving historical metadata. Administrator only.
func (c *Client) ImportWorklog(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueID)+"/worklogs/_import", nil, fields)
}
