package issue

import (
	"context"
	"encoding/json"
)

// LinksListPaginate lists issue links a page at a time, optionally filtered by link type.
func (s *Service) LinksListPaginate(ctx context.Context, issueKey string, perPage, page int, linkTypes []string) (json.RawMessage, error) {
	return s.client.ListIssueLinksPaginate(ctx, issueKey, perPage, page, linkTypes)
}
