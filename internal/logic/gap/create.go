package gap

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CreateSafety = safety.Write

func (s *Service) Create(ctx context.Context, gaps []map[string]any) (json.RawMessage, error) {
	return s.client.CreateGaps(ctx, gaps)
}
