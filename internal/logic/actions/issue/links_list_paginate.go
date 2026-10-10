package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LinksListPaginateSafety() safety.Level { return safety.ReadOnly }

// LinksListPaginate lists issue links a page at a time, optionally filtered by link type.
func (s *Service) LinksListPaginate(ctx context.Context, issueKey string, perPage, page int, linkTypes []string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LinksListPaginateSafety(), func() (json.RawMessage, error) {
		return s.client.ListIssueLinksPaginate(ctx, issueKey, perPage, page, linkTypes)
	})
}
