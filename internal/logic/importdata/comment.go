package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CommentSafety() safety.Level { return safety.Write }

func (s *Service) Comment(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.CommentSafety()); err != nil {
		return nil, err
	}

	return s.client.ImportComment(ctx, issueID, fields)
}
