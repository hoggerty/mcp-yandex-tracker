package report

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) SearchSafety() safety.Level { return safety.ReadOnly }

// Search finds previously generated issue reports.
func (s *Service) Search(ctx context.Context, filter map[string]any, perPage, page int) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.SearchSafety(), func() (json.RawMessage, error) {
		return s.client.SearchReports(ctx, filter, perPage, page)
	})
}
