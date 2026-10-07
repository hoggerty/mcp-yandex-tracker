package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// ListBoards lists the agile boards visible to the current user.
func (c *Client) ListBoards(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/boards", nil, nil)
}

// ListBoardSprints lists sprints that belong to the given board.
func (c *Client) ListBoardSprints(ctx context.Context, boardID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/boards/"+url.PathEscape(boardID)+"/sprints", nil, nil)
}
