package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListQueues lists the queues visible to the current user.
func (c *Client) ListQueues(ctx context.Context, perPage, page int) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	return c.do(ctx, http.MethodGet, "/v2/queues", q, nil)
}

func (c *Client) GetQueue(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/queues/"+url.PathEscape(queueKey), nil, nil)
}
