package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

// Get retrieves a single issue by its key, e.g. "QUEUE-123".
func (s *Service) Get(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.GetIssue(ctx, issueKey)
}
