package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) SearchSafety() safety.Level { return safety.ReadOnly }

// Search finds worklog records across all issues by author and/or creation date range.
func (s *Service) Search(ctx context.Context, createdBy, createdFrom, createdTo string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.SearchSafety(), func() (json.RawMessage, error) {
		return s.client.SearchWorklog(ctx, createdBy, createdFrom, createdTo)
	})
}
