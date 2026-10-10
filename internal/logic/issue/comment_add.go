package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AddCommentSafety() safety.Level { return safety.Write }

func (s *Service) AddComment(ctx context.Context, issueKey, text string, summonees []string) (json.RawMessage, error) {
	if err := s.guard.Check(s.AddCommentSafety()); err != nil {
		return nil, err
	}

	return s.client.AddComment(ctx, issueKey, text, summonees)
}
