package report

import (
	"context"
	"encoding/json"
)

// Create generates an issue report (e.g. an xlsx/csv/xml export of a filtered issue list).
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateReport(ctx, fields)
}
