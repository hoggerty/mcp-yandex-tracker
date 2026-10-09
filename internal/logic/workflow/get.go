package workflow

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

// Get retrieves a single workflow by id.
func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetWorkflow(ctx, id)
}
