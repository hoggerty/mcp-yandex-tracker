package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CommentSafety = safety.Write

func (s *Service) Comment(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.ImportComment(ctx, issueID, fields)
}
