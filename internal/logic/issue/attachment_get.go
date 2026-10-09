package issue

import (
	"context"
	"encoding/json"
)

// AttachmentGet retrieves metadata for a single attached file.
func (s *Service) AttachmentGet(ctx context.Context, issueKey, attachmentID string) (json.RawMessage, error) {
	return s.client.GetAttachmentInfo(ctx, issueKey, attachmentID)
}
