package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
	"github.com/hoggerty/yandex-tracker-mcp/internal/outbound/trackerapi"
)

const SearchSafety = safety.ReadOnly

// Search finds issues using either a query-language string or a structured filter.
func (s *Service) Search(ctx context.Context, query string, filter map[string]any, orderBy string, perPage, page int) (json.RawMessage, error) {
	return s.client.FindIssues(ctx, trackerapi.FindIssuesParams{
		Query:   query,
		Filter:  filter,
		OrderBy: orderBy,
		PerPage: perPage,
		Page:    page,
	})
}
