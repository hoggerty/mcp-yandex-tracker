package sprint

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const StartSafety = safety.Write

// Start changes a sprint's status to "in progress". version is optional (0 to omit).
func (s *Service) Start(ctx context.Context, id string, version int) (json.RawMessage, error) {
	return s.client.StartSprint(ctx, id, version)
}
