package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// CurrentUser retrieves the profile of the user that owns the API token.
func (c *Client) CurrentUser(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v2/myself", nil, nil)
}

// ListUsers lists users known to the organization.
func (c *Client) ListUsers(ctx context.Context, perPage, page int) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	return c.do(ctx, http.MethodGet, "/v2/users", q, nil)
}
