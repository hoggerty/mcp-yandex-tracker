package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListCommentsSafety() safety.Level { return safety.ReadOnly }

func (s *Service) ListComments(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ListCommentsSafety(), func() (json.RawMessage, error) {
		return s.client.ListComments(ctx, issueKey)
	})
}
