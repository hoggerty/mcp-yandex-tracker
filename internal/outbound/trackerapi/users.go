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
	return c.do(ctx, http.MethodGet, "/v3/myself", nil, nil)
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
	return c.do(ctx, http.MethodGet, "/v3/users", q, nil)
}

// GetUser retrieves a single user by login or numeric id.
func (c *Client) GetUser(ctx context.Context, idOrLogin string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/users/"+url.PathEscape(idOrLogin), nil, nil)
}

// GetUsersRelative lists users paginated relative to a starting user id.
func (c *Client) GetUsersRelative(ctx context.Context, perPage, id int, expand string) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if id > 0 {
		q.Set("id", strconv.Itoa(id))
	}
	if expand != "" {
		q.Set("expand", expand)
	}
	return c.do(ctx, http.MethodGet, "/v3/users/_relative", q, nil)
}
