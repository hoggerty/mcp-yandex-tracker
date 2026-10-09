package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListColumns lists the columns of an agile board.
func (c *Client) ListColumns(ctx context.Context, boardID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/boards/"+url.PathEscape(boardID)+"/columns", nil, nil)
}

// GetColumn retrieves the parameters of a single board column.
func (c *Client) GetColumn(ctx context.Context, boardID, columnID string) (json.RawMessage, error) {
	path := "/v3/boards/" + url.PathEscape(boardID) + "/columns/" + url.PathEscape(columnID)
	return c.do(ctx, http.MethodGet, path, nil, nil)
}

// CreateColumn creates a new column on an agile board; version must match the board's current version.
func (c *Client) CreateColumn(ctx context.Context, boardID string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPost, "/v3/boards/"+url.PathEscape(boardID)+"/columns", q, fields)
}

// EditColumn edits a board column; version must match the board's current version.
func (c *Client) EditColumn(ctx context.Context, boardID, columnID string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	path := "/v3/boards/" + url.PathEscape(boardID) + "/columns/" + url.PathEscape(columnID)
	return c.do(ctx, http.MethodPatch, path, q, fields)
}

// DeleteColumn deletes a board column; version must match the board's current version.
func (c *Client) DeleteColumn(ctx context.Context, boardID, columnID string, version int) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	path := "/v3/boards/" + url.PathEscape(boardID) + "/columns/" + url.PathEscape(columnID)
	return c.do(ctx, http.MethodDelete, path, q, nil)
}
