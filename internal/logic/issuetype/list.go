package issuetype

import (
	"context"
	"encoding/json"
)

// List lists the issue types available in the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListIssueTypes(ctx)
}
