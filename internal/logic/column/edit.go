package column

import (
	"context"
	"encoding/json"
)

// Edit edits a board column; version must match the board's current version.
func (s *Service) Edit(ctx context.Context, boardID, columnID string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditColumn(ctx, boardID, columnID, version, fields)
}
