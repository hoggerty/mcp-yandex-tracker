package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const SearchReleaseSafety = safety.Write

// SearchRelease releases server-side resources held by a scrolled search-issues request.
func (s *Service) SearchRelease(ctx context.Context, scrollID, scrollToken string) (json.RawMessage, error) {
	return s.client.ReleaseSearchScroll(ctx, scrollID, scrollToken)
}
