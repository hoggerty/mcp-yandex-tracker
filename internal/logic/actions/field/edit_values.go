package field

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) EditValuesSafety() safety.Level { return safety.Write }

// EditValues edits a global field's non-name parameters (category, order, description,
// visibility, possible values); version must match the field's current version.
func (s *Service) EditValues(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.EditValuesSafety(), func() (json.RawMessage, error) {
		return s.client.PatchField(ctx, id, version, fields)
	})
}
