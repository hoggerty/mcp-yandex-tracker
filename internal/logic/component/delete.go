package component

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

// Delete deletes a component.
func (s *Service) Delete(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.DeleteComponent(ctx, id)
}
