package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// ListApplications lists external applications known to Tracker (used as link origins).
func (c *Client) ListApplications(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/applications", nil, nil)
}

// ListExternalLinks lists an issue's links to objects in external applications.
func (c *Client) ListExternalLinks(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/remotelinks", nil, nil)
}

// AddExternalLink links an issue to an object in an external application.
func (c *Client) AddExternalLink(ctx context.Context, issueKey, relationship, key, origin string) (json.RawMessage, error) {
	body := map[string]any{"relationship": relationship, "key": key, "origin": origin}
	q := url.Values{}
	q.Set("backlink", "true")
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/remotelinks", q, body)
}

// DeleteExternalLink removes an issue's link to an external application object.
func (c *Client) DeleteExternalLink(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/issues/"+url.PathEscape(issueKey)+"/remotelinks/"+url.PathEscape(linkID), nil, nil)
}
