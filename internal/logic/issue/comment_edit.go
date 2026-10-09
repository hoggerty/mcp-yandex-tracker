package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CommentEditSafety = safety.Write

// CommentEdit changes the text (and optionally summonees) of an existing comment.
func (s *Service) CommentEdit(ctx context.Context, issueKey, commentID, text string, summonees []string) (json.RawMessage, error) {
	return s.client.EditComment(ctx, issueKey, commentID, text, summonees)
}
