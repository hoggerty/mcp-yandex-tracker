package column

import (
	"context"
	"encoding/json"
)

// Get retrieves the parameters of a single board column.
func (s *Service) Get(ctx context.Context, boardID, columnID string) (json.RawMessage, error) {
	return s.client.GetColumn(ctx, boardID, columnID)
}
