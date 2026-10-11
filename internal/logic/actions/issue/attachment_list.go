package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) AttachmentListSafety() safety.Level { return safety.ReadOnly }

// AttachmentList lists the files attached to an issue.
func (s *Service) AttachmentList(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AttachmentListSafety(), func() (json.RawMessage, error) {
		return s.client.ListAttachments(ctx, issueKey)
	})
}
