package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListCommentsSafety = safety.ReadOnly

func (s *Service) ListComments(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListComments(ctx, issueKey)
}
