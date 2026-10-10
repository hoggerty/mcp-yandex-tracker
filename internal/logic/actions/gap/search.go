package gap

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) SearchSafety() safety.Level { return safety.ReadOnly }

func (s *Service) Search(ctx context.Context, body map[string]any, perPage, page int) (json.RawMessage, error) {
	if err := s.guard.Check(s.SearchSafety()); err != nil {
		return nil, err
	}

	return s.client.SearchGaps(ctx, body, perPage, page)
}
