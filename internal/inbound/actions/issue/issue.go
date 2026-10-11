// Package issue exposes Yandex Tracker issue operations as MCP tools.
package issue

import (
	"github.com/hoggerty/mcp-yandex-tracker/internal/inbound/tool"
	logicissue "github.com/hoggerty/mcp-yandex-tracker/internal/logic/actions/issue"
)

func Tools(svc *logicissue.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &GetTool{Svc: svc}, Safety: svc.GetSafety()},
		{Tool: &SearchTool{Svc: svc}, Safety: svc.SearchSafety()},
		{Tool: &CountTool{Svc: svc}, Safety: svc.CountSafety()},
		{Tool: &CreateTool{Svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &UpdateTool{Svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &TransitionsListTool{Svc: svc}, Safety: svc.ListTransitionsSafety()},
		{Tool: &TransitionExecuteTool{Svc: svc}, Safety: svc.ExecuteTransitionSafety()},
		{Tool: &CommentsListTool{Svc: svc}, Safety: svc.ListCommentsSafety()},
		{Tool: &CommentAddTool{Svc: svc}, Safety: svc.AddCommentSafety()},
		{Tool: &LinksListTool{Svc: svc}, Safety: svc.ListLinksSafety()},
		{Tool: &LinkAddTool{Svc: svc}, Safety: svc.LinkIssuesSafety()},
		{Tool: &ChangelogGetTool{Svc: svc}, Safety: svc.GetChangelogSafety()},
		{Tool: &MoveTool{Svc: svc}, Safety: svc.MoveSafety()},
		{Tool: &SuggestTool{Svc: svc}, Safety: svc.SuggestSafety()},
		{Tool: &SearchReleaseTool{Svc: svc}, Safety: svc.SearchReleaseSafety()},
		{Tool: &ChecklistAddTool{Svc: svc}, Safety: svc.ChecklistAddSafety()},
		{Tool: &ChecklistListTool{Svc: svc}, Safety: svc.ChecklistListSafety()},
		{Tool: &ChecklistEditTool{Svc: svc}, Safety: svc.ChecklistEditSafety()},
		{Tool: &ChecklistDeleteTool{Svc: svc}, Safety: svc.ChecklistDeleteSafety()},
		{Tool: &ChecklistItemDeleteTool{Svc: svc}, Safety: svc.ChecklistItemDeleteSafety()},
		{Tool: &CommentGetTool{Svc: svc}, Safety: svc.CommentGetSafety()},
		{Tool: &CommentEditTool{Svc: svc}, Safety: svc.CommentEditSafety()},
		{Tool: &CommentDeleteTool{Svc: svc}, Safety: svc.CommentDeleteSafety()},
		{Tool: &CommentReactionAddTool{Svc: svc}, Safety: svc.CommentReactionAddSafety()},
		{Tool: &LinksListPaginateTool{Svc: svc}, Safety: svc.LinksListPaginateSafety()},
		{Tool: &LinkDeleteTool{Svc: svc}, Safety: svc.LinkDeleteSafety()},
		{Tool: &AttachmentListTool{Svc: svc}, Safety: svc.AttachmentListSafety()},
		{Tool: &AttachmentGetTool{Svc: svc}, Safety: svc.AttachmentGetSafety()},
	}
}
