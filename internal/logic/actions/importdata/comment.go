package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) CommentSafety() safety.Level { return safety.Write }

func (s *Service) Comment(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.CommentSafety(), func() (json.RawMessage, error) {
		return s.client.ImportComment(ctx, issueID, fields)
	})
}
