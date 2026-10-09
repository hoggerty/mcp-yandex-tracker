package field

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CreateSafety = safety.Write

// Create creates a new global issue field.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateField(ctx, fields)
}
