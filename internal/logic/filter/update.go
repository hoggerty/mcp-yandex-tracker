package filter

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) UpdateSafety() safety.Level { return safety.Write }

// Update replaces a saved issue filter's name/filter/sorts.
func (s *Service) Update(ctx context.Context, id string, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.UpdateSafety()); err != nil {
		return nil, err
	}

	return s.client.UpdateFilter(ctx, id, fields)
}
