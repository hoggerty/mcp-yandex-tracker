package gap

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

func (s *Service) Delete(ctx context.Context, gapIDs []string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.DeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteGaps(ctx, gapIDs)
	})
}
