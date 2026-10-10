package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) SearchReleaseSafety() safety.Level { return safety.Write }

// SearchRelease releases server-side resources held by a scrolled search-issues request.
func (s *Service) SearchRelease(ctx context.Context, scrollID, scrollToken string) (json.RawMessage, error) {
	if err := s.guard.Check(s.SearchReleaseSafety()); err != nil {
		return nil, err
	}

	return s.client.ReleaseSearchScroll(ctx, scrollID, scrollToken)
}
