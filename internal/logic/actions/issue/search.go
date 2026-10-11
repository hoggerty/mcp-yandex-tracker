package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
	"github.com/hoggerty/mcp-yandex-tracker/internal/outbound/trackerapi"
)

func (s *Service) SearchSafety() safety.Level { return safety.ReadOnly }

// Search finds issues using either a query-language string or a structured filter.
func (s *Service) Search(ctx context.Context, query string, filter map[string]any, orderBy string, perPage, page int) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.SearchSafety(), func() (json.RawMessage, error) {
		return s.client.FindIssues(ctx, trackerapi.FindIssuesParams{
			Query:   query,
			Filter:  filter,
			OrderBy: orderBy,
			PerPage: perPage,
			Page:    page,
		})
	})
}
