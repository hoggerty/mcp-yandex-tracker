package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) SearchSafety() safety.Level { return safety.ReadOnly }

// Search finds worklog records across all issues by author and/or creation date range.
func (s *Service) Search(ctx context.Context, createdBy, createdFrom, createdTo string) (json.RawMessage, error) {
	if err := s.guard.Check(s.SearchSafety()); err != nil {
		return nil, err
	}

	return s.client.SearchWorklog(ctx, createdBy, createdFrom, createdTo)
}
