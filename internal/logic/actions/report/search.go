package report

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) SearchSafety() safety.Level { return safety.ReadOnly }

// Search finds previously generated issue reports.
func (s *Service) Search(ctx context.Context, filter map[string]any, perPage, page int) (json.RawMessage, error) {
	if err := s.guard.Check(s.SearchSafety()); err != nil {
		return nil, err
	}

	return s.client.SearchReports(ctx, filter, perPage, page)
}
