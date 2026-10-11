package sprint

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ArchiveSafety() safety.Level { return safety.Destructive }

// Archive archives a sprint. version is optional (0 to omit).
func (s *Service) Archive(ctx context.Context, id string, version int) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ArchiveSafety(), func() (json.RawMessage, error) {
		return s.client.ArchiveSprint(ctx, id, version)
	})
}
