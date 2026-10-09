package field

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CategoryCreateSafety = safety.Write

// CategoryCreate creates a new category for grouping global fields.
func (s *Service) CategoryCreate(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateFieldCategory(ctx, fields)
}
