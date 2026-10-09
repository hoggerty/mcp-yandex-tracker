package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// AddWorklog records time spent on an issue.
func (c *Client) AddWorklog(ctx context.Context, issueKey, start, duration, comment string) (json.RawMessage, error) {
	body := map[string]any{"start": start, "duration": duration}
	if comment != "" {
		body["comment"] = comment
	}
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/worklog", nil, body)
}

// EditWorklog changes the duration and/or comment of an existing worklog record.
func (c *Client) EditWorklog(ctx context.Context, issueKey, worklogID, duration, comment string) (json.RawMessage, error) {
	body := map[string]any{"duration": duration}
	if comment != "" {
		body["comment"] = comment
	}
	return c.do(ctx, http.MethodPatch, "/v3/issues/"+url.PathEscape(issueKey)+"/worklog/"+url.PathEscape(worklogID), nil, body)
}

// DeleteWorklog deletes a worklog record from an issue.
func (c *Client) DeleteWorklog(ctx context.Context, issueKey, worklogID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/issues/"+url.PathEscape(issueKey)+"/worklog/"+url.PathEscape(worklogID), nil, nil)
}

// ListWorklog lists all worklog records for an issue.
func (c *Client) ListWorklog(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/worklog", nil, nil)
}

// SearchWorklog searches worklog records across all issues by author and/or creation date range.
func (c *Client) SearchWorklog(ctx context.Context, createdBy, createdFrom, createdTo string) (json.RawMessage, error) {
	body := map[string]any{}
	if createdBy != "" {
		body["createdBy"] = createdBy
	}
	if createdFrom != "" || createdTo != "" {
		createdAt := map[string]any{}
		if createdFrom != "" {
			createdAt["from"] = createdFrom
		}
		if createdTo != "" {
			createdAt["to"] = createdTo
		}
		body["createdAt"] = createdAt
	}
	return c.do(ctx, http.MethodPost, "/v3/worklog/_search", nil, body)
}
