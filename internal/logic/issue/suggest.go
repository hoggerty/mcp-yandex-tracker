package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const SuggestSafety = safety.ReadOnly

// Suggest returns issue suggestions matching input text, optionally scoped to a queue.
func (s *Service) Suggest(ctx context.Context, input, queue string) (json.RawMessage, error) {
	return s.client.SuggestIssues(ctx, input, queue)
}
