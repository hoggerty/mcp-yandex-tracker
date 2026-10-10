package gap

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) SearchSafety() safety.Level { return safety.ReadOnly }

func (s *Service) Search(ctx context.Context, body map[string]any, perPage, page int) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.SearchSafety(), func() (json.RawMessage, error) {
		return s.client.SearchGaps(ctx, body, perPage, page)
	})
}
