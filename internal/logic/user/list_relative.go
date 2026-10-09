package user

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListRelativeSafety = safety.ReadOnly

// ListRelative lists users paginated relative to a starting user id.
func (s *Service) ListRelative(ctx context.Context, perPage, id int, expand string) (json.RawMessage, error) {
	return s.client.GetUsersRelative(ctx, perPage, id, expand)
}
