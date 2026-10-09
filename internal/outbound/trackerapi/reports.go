package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// CreateReport generates an issue report (e.g. an xlsx/csv/xml export of a filtered issue list).
func (c *Client) CreateReport(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/entities/report/", nil, fields)
}

// SearchReports finds previously generated issue reports.
func (c *Client) SearchReports(ctx context.Context, filter map[string]any, perPage, page int) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	body := map[string]any{}
	if len(filter) > 0 {
		body["filter"] = filter
	}
	return c.do(ctx, http.MethodPost, "/v3/entities/report/_search", q, body)
}
