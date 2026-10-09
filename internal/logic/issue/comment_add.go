package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AddCommentSafety = safety.Write

func (s *Service) AddComment(ctx context.Context, issueKey, text string, summonees []string) (json.RawMessage, error) {
	return s.client.AddComment(ctx, issueKey, text, summonees)
}
