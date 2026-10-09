package version

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const EditSafety = safety.Write

// Edit edits a version.
func (s *Service) Edit(ctx context.Context, id string, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditVersion(ctx, id, fields)
}
