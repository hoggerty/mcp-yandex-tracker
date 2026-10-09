package sprint

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ArchiveSafety = safety.Destructive

// Archive archives a sprint. version is optional (0 to omit).
func (s *Service) Archive(ctx context.Context, id string, version int) (json.RawMessage, error) {
	return s.client.ArchiveSprint(ctx, id, version)
}
