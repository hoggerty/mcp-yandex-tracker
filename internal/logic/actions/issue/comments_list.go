package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListCommentsSafety() safety.Level { return safety.ReadOnly }

func (s *Service) ListComments(ctx context.Context, issueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListCommentsSafety()); err != nil {
		return nil, err
	}

	return s.client.ListComments(ctx, issueKey)
}
