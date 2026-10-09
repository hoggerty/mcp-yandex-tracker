// Package issue exposes Yandex Tracker issue operations as MCP tools.
package issue

import (
	"github.com/hoggerty/yandex-tracker-mcp/internal/inbound/tool"
	logicissue "github.com/hoggerty/yandex-tracker-mcp/internal/logic/issue"
)

func Tools(svc *logicissue.Service) []tool.Entry {
	return []tool.Entry{
		{Tool: &getTool{svc: svc}, Safety: logicissue.GetSafety},
		{Tool: &searchTool{svc: svc}, Safety: logicissue.SearchSafety},
		{Tool: &countTool{svc: svc}, Safety: logicissue.CountSafety},
		{Tool: &createTool{svc: svc}, Safety: logicissue.CreateSafety},
		{Tool: &updateTool{svc: svc}, Safety: logicissue.UpdateSafety},
		{Tool: &transitionsListTool{svc: svc}, Safety: logicissue.ListTransitionsSafety},
		{Tool: &transitionExecuteTool{svc: svc}, Safety: logicissue.ExecuteTransitionSafety},
		{Tool: &commentsListTool{svc: svc}, Safety: logicissue.ListCommentsSafety},
		{Tool: &commentAddTool{svc: svc}, Safety: logicissue.AddCommentSafety},
		{Tool: &linksListTool{svc: svc}, Safety: logicissue.ListLinksSafety},
		{Tool: &linkAddTool{svc: svc}, Safety: logicissue.LinkIssuesSafety},
		{Tool: &changelogGetTool{svc: svc}, Safety: logicissue.GetChangelogSafety},
		{Tool: &moveTool{svc: svc}, Safety: logicissue.MoveSafety},
		{Tool: &suggestTool{svc: svc}, Safety: logicissue.SuggestSafety},
		{Tool: &searchReleaseTool{svc: svc}, Safety: logicissue.SearchReleaseSafety},
		{Tool: &checklistAddTool{svc: svc}, Safety: logicissue.ChecklistAddSafety},
		{Tool: &checklistListTool{svc: svc}, Safety: logicissue.ChecklistListSafety},
		{Tool: &checklistEditTool{svc: svc}, Safety: logicissue.ChecklistEditSafety},
		{Tool: &checklistDeleteTool{svc: svc}, Safety: logicissue.ChecklistDeleteSafety},
		{Tool: &checklistItemDeleteTool{svc: svc}, Safety: logicissue.ChecklistItemDeleteSafety},
		{Tool: &commentGetTool{svc: svc}, Safety: logicissue.CommentGetSafety},
		{Tool: &commentEditTool{svc: svc}, Safety: logicissue.CommentEditSafety},
		{Tool: &commentDeleteTool{svc: svc}, Safety: logicissue.CommentDeleteSafety},
		{Tool: &commentReactionAddTool{svc: svc}, Safety: logicissue.CommentReactionAddSafety},
		{Tool: &linksListPaginateTool{svc: svc}, Safety: logicissue.LinksListPaginateSafety},
		{Tool: &linkDeleteTool{svc: svc}, Safety: logicissue.LinkDeleteSafety},
		{Tool: &attachmentListTool{svc: svc}, Safety: logicissue.AttachmentListSafety},
		{Tool: &attachmentGetTool{svc: svc}, Safety: logicissue.AttachmentGetSafety},
	}
}
