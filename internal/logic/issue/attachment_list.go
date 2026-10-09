package issue

import (
	"context"
	"encoding/json"
)

// AttachmentList lists the files attached to an issue.
func (s *Service) AttachmentList(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListAttachments(ctx, issueKey)
}
