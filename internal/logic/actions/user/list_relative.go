package user

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ListRelativeSafety() safety.Level { return safety.ReadOnly }

// ListRelative lists users paginated relative to a starting user id.
func (s *Service) ListRelative(ctx context.Context, perPage, id int, expand string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ListRelativeSafety(), func() (json.RawMessage, error) {
		return s.client.GetUsersRelative(ctx, perPage, id, expand)
	})
}
