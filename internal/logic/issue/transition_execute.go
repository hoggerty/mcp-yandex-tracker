package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ExecuteTransitionSafety = safety.Write

func (s *Service) ExecuteTransition(ctx context.Context, issueKey, transitionID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.ExecuteTransition(ctx, issueKey, transitionID, fields)
}
