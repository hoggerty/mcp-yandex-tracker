package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type FindIssuesParams struct {
	Query   string
	Filter  map[string]any
	OrderBy string
	PerPage int
	Page    int
}

// GetIssue retrieves a single issue by its key, e.g. "QUEUE-123".
func (c *Client) GetIssue(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/issues/"+url.PathEscape(issueKey), nil, nil)
}

// FindIssues searches for issues using either a Tracker query language string
// or a structured filter.
func (c *Client) FindIssues(ctx context.Context, p FindIssuesParams) (json.RawMessage, error) {
	body := map[string]any{}
	if p.Query != "" {
		body["query"] = p.Query
	}
	if len(p.Filter) > 0 {
		body["filter"] = p.Filter
	}
	if p.OrderBy != "" {
		body["order"] = p.OrderBy
	}

	q := url.Values{}
	if p.PerPage > 0 {
		q.Set("perPage", strconv.Itoa(p.PerPage))
	}
	if p.Page > 0 {
		q.Set("page", strconv.Itoa(p.Page))
	}

	return c.do(ctx, http.MethodPost, "/v2/issues/_search", q, body)
}

func (c *Client) CountIssues(ctx context.Context, query string, filter map[string]any) (json.RawMessage, error) {
	body := map[string]any{}
	if query != "" {
		body["query"] = query
	}
	if len(filter) > 0 {
		body["filter"] = filter
	}
	return c.do(ctx, http.MethodPost, "/v2/issues/_count", nil, body)
}

// CreateIssue creates a new issue from a field map, e.g.
// {"summary": "...", "queue": "QUEUE", "description": "..."}.
func (c *Client) CreateIssue(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v2/issues", nil, fields)
}

func (c *Client) UpdateIssue(ctx context.Context, issueKey string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPatch, "/v2/issues/"+url.PathEscape(issueKey), nil, fields)
}

func (c *Client) ListTransitions(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/issues/"+url.PathEscape(issueKey)+"/transitions", nil, nil)
}

func (c *Client) ExecuteTransition(ctx context.Context, issueKey, transitionID string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v2/issues/"+url.PathEscape(issueKey)+"/transitions/"+url.PathEscape(transitionID)+"/_execute", nil, fields)
}

func (c *Client) ListComments(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/issues/"+url.PathEscape(issueKey)+"/comments", nil, nil)
}

func (c *Client) AddComment(ctx context.Context, issueKey, text string, summonees []string) (json.RawMessage, error) {
	body := map[string]any{"text": text}
	if len(summonees) > 0 {
		body["summonees"] = summonees
	}
	return c.do(ctx, http.MethodPost, "/v2/issues/"+url.PathEscape(issueKey)+"/comments", nil, body)
}

// ListIssueLinks lists the issue links (relates to, blocks, etc.) for an issue.
func (c *Client) ListIssueLinks(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/issues/"+url.PathEscape(issueKey)+"/links", nil, nil)
}

// LinkIssues creates a link between two issues, e.g. relationship "relates" or "is dependent by".
func (c *Client) LinkIssues(ctx context.Context, issueKey, relationship, targetIssueKey string) (json.RawMessage, error) {
	body := map[string]any{
		"relationship": relationship,
		"issue":        targetIssueKey,
	}
	return c.do(ctx, http.MethodPost, "/v2/issues/"+url.PathEscape(issueKey)+"/links", nil, body)
}

// GetChangelog retrieves the field change history for an issue.
func (c *Client) GetChangelog(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/issues/"+url.PathEscape(issueKey)+"/changelog", nil, nil)
}
