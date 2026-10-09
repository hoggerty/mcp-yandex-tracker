package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CreateGaps creates one or more employee absence records. Administrator only.
func (c *Client) CreateGaps(ctx context.Context, gaps []map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/gaps", nil, map[string]any{"gaps": gaps})
}

// SearchGaps searches employee absences for the given users and date range. Administrator only.
func (c *Client) SearchGaps(ctx context.Context, body map[string]any, perPage, page int) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	return c.do(ctx, http.MethodPost, "/v3/gaps/_search", q, body)
}

// DeleteGaps deletes employee absence records by id. Administrator only.
func (c *Client) DeleteGaps(ctx context.Context, gapIDs []string) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("gapIds", strings.Join(gapIDs, ","))
	return c.do(ctx, http.MethodDelete, "/v3/gaps", q, nil)
}
