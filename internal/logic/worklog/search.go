package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const SearchSafety = safety.ReadOnly

// Search finds worklog records across all issues by author and/or creation date range.
func (s *Service) Search(ctx context.Context, createdBy, createdFrom, createdTo string) (json.RawMessage, error) {
	return s.client.SearchWorklog(ctx, createdBy, createdFrom, createdTo)
}
