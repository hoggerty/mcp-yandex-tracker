package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LinkSafety() safety.Level { return safety.Write }

func (s *Service) Link(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LinkSafety(), func() (json.RawMessage, error) {
		return s.client.ImportLink(ctx, issueID, fields)
	})
}
