package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AttachmentListSafety() safety.Level { return safety.ReadOnly }

// AttachmentList lists the files attached to an issue.
func (s *Service) AttachmentList(ctx context.Context, issueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.AttachmentListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListAttachments(ctx, issueKey)
}
