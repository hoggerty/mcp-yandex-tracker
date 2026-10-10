package report

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CreateSafety() safety.Level { return safety.Write }

// Create generates an issue report (e.g. an xlsx/csv/xml export of a filtered issue list).
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.CreateSafety()); err != nil {
		return nil, err
	}

	return s.client.CreateReport(ctx, fields)
}
