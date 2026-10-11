package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) SuggestSafety() safety.Level { return safety.ReadOnly }

// Suggest returns issue suggestions matching input text, optionally scoped to a queue.
func (s *Service) Suggest(ctx context.Context, input, queue string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.SuggestSafety(), func() (json.RawMessage, error) {
		return s.client.SuggestIssues(ctx, input, queue)
	})
}
