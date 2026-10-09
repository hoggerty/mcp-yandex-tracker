package column

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

// Get retrieves the parameters of a single board column.
func (s *Service) Get(ctx context.Context, boardID, columnID string) (json.RawMessage, error) {
	return s.client.GetColumn(ctx, boardID, columnID)
}
