package issue

import (
	"context"
	"encoding/json"
)

// ChecklistDelete deletes an issue's entire checklist.
func (s *Service) ChecklistDelete(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.DeleteChecklist(ctx, issueKey)
}
