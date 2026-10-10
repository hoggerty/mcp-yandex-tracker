package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) SuggestSafety() safety.Level { return safety.ReadOnly }

// Suggest returns issue suggestions matching input text, optionally scoped to a queue.
func (s *Service) Suggest(ctx context.Context, input, queue string) (json.RawMessage, error) {
	if err := s.guard.Check(s.SuggestSafety()); err != nil {
		return nil, err
	}

	return s.client.SuggestIssues(ctx, input, queue)
}
