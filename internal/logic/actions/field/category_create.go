package field

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CategoryCreateSafety() safety.Level { return safety.Write }

// CategoryCreate creates a new category for grouping global fields.
func (s *Service) CategoryCreate(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.CategoryCreateSafety()); err != nil {
		return nil, err
	}

	return s.client.CreateFieldCategory(ctx, fields)
}
