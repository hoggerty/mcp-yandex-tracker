package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const EditSafety = safety.Write

// Edit edits an entity comment.
func (s *Service) Edit(ctx context.Context, entityType, id, commentID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditEntityComment(ctx, entityType, id, commentID, fields)
}
