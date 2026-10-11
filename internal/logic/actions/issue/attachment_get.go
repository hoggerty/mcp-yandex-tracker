package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) AttachmentGetSafety() safety.Level { return safety.ReadOnly }

// AttachmentGet retrieves metadata for a single attached file.
func (s *Service) AttachmentGet(ctx context.Context, issueKey, attachmentID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AttachmentGetSafety(), func() (json.RawMessage, error) {
		return s.client.GetAttachmentInfo(ctx, issueKey, attachmentID)
	})
}
