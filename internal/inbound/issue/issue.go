// Package issue exposes Yandex Tracker issue operations as MCP tools.
package issue

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

func Tools(svc *logicissue.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &getTool{svc: svc}, Safety: svc.GetSafety()},
		{Tool: &searchTool{svc: svc}, Safety: svc.SearchSafety()},
		{Tool: &countTool{svc: svc}, Safety: svc.CountSafety()},
		{Tool: &createTool{svc: svc}, Safety: svc.CreateSafety()},
		{Tool: &updateTool{svc: svc}, Safety: svc.UpdateSafety()},
		{Tool: &transitionsListTool{svc: svc}, Safety: svc.ListTransitionsSafety()},
		{Tool: &transitionExecuteTool{svc: svc}, Safety: svc.ExecuteTransitionSafety()},
		{Tool: &commentsListTool{svc: svc}, Safety: svc.ListCommentsSafety()},
		{Tool: &commentAddTool{svc: svc}, Safety: svc.AddCommentSafety()},
		{Tool: &linksListTool{svc: svc}, Safety: svc.ListLinksSafety()},
		{Tool: &linkAddTool{svc: svc}, Safety: svc.LinkIssuesSafety()},
		{Tool: &changelogGetTool{svc: svc}, Safety: svc.GetChangelogSafety()},
		{Tool: &moveTool{svc: svc}, Safety: svc.MoveSafety()},
		{Tool: &suggestTool{svc: svc}, Safety: svc.SuggestSafety()},
		{Tool: &searchReleaseTool{svc: svc}, Safety: svc.SearchReleaseSafety()},
		{Tool: &checklistAddTool{svc: svc}, Safety: svc.ChecklistAddSafety()},
		{Tool: &checklistListTool{svc: svc}, Safety: svc.ChecklistListSafety()},
		{Tool: &checklistEditTool{svc: svc}, Safety: svc.ChecklistEditSafety()},
		{Tool: &checklistDeleteTool{svc: svc}, Safety: svc.ChecklistDeleteSafety()},
		{Tool: &checklistItemDeleteTool{svc: svc}, Safety: svc.ChecklistItemDeleteSafety()},
		{Tool: &commentGetTool{svc: svc}, Safety: svc.CommentGetSafety()},
		{Tool: &commentEditTool{svc: svc}, Safety: svc.CommentEditSafety()},
		{Tool: &commentDeleteTool{svc: svc}, Safety: svc.CommentDeleteSafety()},
		{Tool: &commentReactionAddTool{svc: svc}, Safety: svc.CommentReactionAddSafety()},
		{Tool: &linksListPaginateTool{svc: svc}, Safety: svc.LinksListPaginateSafety()},
		{Tool: &linkDeleteTool{svc: svc}, Safety: svc.LinkDeleteSafety()},
		{Tool: &attachmentListTool{svc: svc}, Safety: svc.AttachmentListSafety()},
		{Tool: &attachmentGetTool{svc: svc}, Safety: svc.AttachmentGetSafety()},
	}
}
