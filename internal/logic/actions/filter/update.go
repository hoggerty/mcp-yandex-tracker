package filter

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) UpdateSafety() safety.Level { return safety.Write }

// Update replaces a saved issue filter's name/filter/sorts.
func (s *Service) Update(ctx context.Context, id string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.UpdateSafety(), func() (json.RawMessage, error) {
		return s.client.UpdateFilter(ctx, id, fields)
	})
}
