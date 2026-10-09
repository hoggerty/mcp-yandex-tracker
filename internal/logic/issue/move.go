package issue

import (
	"context"
	"encoding/json"
)

// Move moves an issue to another queue.
func (s *Service) Move(ctx context.Context, issueKey, targetQueue string, notify, notifyAuthor, moveAllFields, initialStatus bool, fields map[string]any) (json.RawMessage, error) {
	return s.client.MoveIssue(ctx, issueKey, targetQueue, notify, notifyAuthor, moveAllFields, initialStatus, fields)
}
