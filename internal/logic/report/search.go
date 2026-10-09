package report

import (
	"context"
	"encoding/json"
)

// Search finds previously generated issue reports.
func (s *Service) Search(ctx context.Context, filter map[string]any, perPage, page int) (json.RawMessage, error) {
	return s.client.SearchReports(ctx, filter, perPage, page)
}
