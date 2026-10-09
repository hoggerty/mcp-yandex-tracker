package gap

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

func (s *Service) Delete(ctx context.Context, gapIDs []string) (json.RawMessage, error) {
	return s.client.DeleteGaps(ctx, gapIDs)
}
