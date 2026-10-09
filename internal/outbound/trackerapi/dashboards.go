package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// CreateDashboard creates a new dashboard.
func (c *Client) CreateDashboard(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/dashboards", nil, fields)
}

// CreateCycleTimeWidget creates a "Cycle time" widget on a dashboard.
func (c *Client) CreateCycleTimeWidget(ctx context.Context, dashboardID string, fields map[string]any) (json.RawMessage, error) {
	path := "/v3/dashboards/" + url.PathEscape(dashboardID) + "/widgets/cycleTime"
	return c.do(ctx, http.MethodPost, path, nil, fields)
}
