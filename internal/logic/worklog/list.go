package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List returns all worklog records for an issue.
func (s *Service) List(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListWorklog(ctx, issueKey)
}
