package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListBoards lists the agile boards visible to the current user, sorted by
// ID; pass the last board's ID from the previous page as id to continue.
func (c *Client) ListBoards(ctx context.Context, perPage, id int) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if id > 0 {
		q.Set("id", strconv.Itoa(id))
	}
	return c.do(ctx, http.MethodGet, "/v3/boards/_paginate", q, nil)
}

func (c *Client) ListBoardSprints(ctx context.Context, boardID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/boards/"+url.PathEscape(boardID)+"/sprints", nil, nil)
}

// GetBoard retrieves the parameters of a single agile board.
func (c *Client) GetBoard(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/boards/"+url.PathEscape(id), nil, nil)
}

// CreateBoard creates a new agile board. Uses the current, non-deprecated
// /v3/liveBoards endpoint (plain POST /v3/boards is documented as deprecated).
func (c *Client) CreateBoard(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/liveBoards", nil, fields)
}

// EditBoard edits an agile board; version must match the board's current version.
func (c *Client) EditBoard(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/boards/"+url.PathEscape(id), q, fields)
}

// DeleteBoard deletes an agile board.
func (c *Client) DeleteBoard(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/boards/"+url.PathEscape(id), nil, nil)
}
