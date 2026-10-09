package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type FindIssuesParams struct {
	Query   string
	Filter  map[string]any
	OrderBy string
	PerPage int
	Page    int
}

// GetIssue retrieves a single issue by its key, e.g. "QUEUE-123".
func (c *Client) GetIssue(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey), nil, nil)
}

// FindIssues searches for issues using either a Tracker query language string
// or a structured filter.
func (c *Client) FindIssues(ctx context.Context, p FindIssuesParams) (json.RawMessage, error) {
	body := map[string]any{}
	if p.Query != "" {
		body["query"] = p.Query
	}
	if len(p.Filter) > 0 {
		body["filter"] = p.Filter
	}
	if p.OrderBy != "" {
		body["order"] = p.OrderBy
	}

	q := url.Values{}
	if p.PerPage > 0 {
		q.Set("perPage", strconv.Itoa(p.PerPage))
	}
	if p.Page > 0 {
		q.Set("page", strconv.Itoa(p.Page))
	}

	return c.do(ctx, http.MethodPost, "/v3/issues/_search", q, body)
}

func (c *Client) CountIssues(ctx context.Context, query string, filter map[string]any) (json.RawMessage, error) {
	body := map[string]any{}
	if query != "" {
		body["query"] = query
	}
	if len(filter) > 0 {
		body["filter"] = filter
	}
	return c.do(ctx, http.MethodPost, "/v3/issues/_count", nil, body)
}

// CreateIssue creates a new issue from a field map, e.g.
// {"summary": "...", "queue": "QUEUE", "description": "..."}.
func (c *Client) CreateIssue(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issues", nil, fields)
}

func (c *Client) UpdateIssue(ctx context.Context, issueKey string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPatch, "/v3/issues/"+url.PathEscape(issueKey), nil, fields)
}

func (c *Client) ListTransitions(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/transitions", nil, nil)
}

func (c *Client) ExecuteTransition(ctx context.Context, issueKey, transitionID string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/transitions/"+url.PathEscape(transitionID)+"/_execute", nil, fields)
}

func (c *Client) ListComments(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/comments", nil, nil)
}

func (c *Client) AddComment(ctx context.Context, issueKey, text string, summonees []string) (json.RawMessage, error) {
	body := map[string]any{"text": text}
	if len(summonees) > 0 {
		body["summonees"] = summonees
	}
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/comments", nil, body)
}

// ListIssueLinks lists the issue links (relates to, blocks, etc.) for an issue.
func (c *Client) ListIssueLinks(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/links", nil, nil)
}

// LinkIssues creates a link between two issues, e.g. relationship "relates" or "is dependent by".
func (c *Client) LinkIssues(ctx context.Context, issueKey, relationship, targetIssueKey string) (json.RawMessage, error) {
	body := map[string]any{
		"relationship": relationship,
		"issue":        targetIssueKey,
	}
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/links", nil, body)
}

// GetChangelog retrieves the field change history for an issue.
func (c *Client) GetChangelog(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/changelog", nil, nil)
}

// MoveIssue moves an issue to another queue.
func (c *Client) MoveIssue(ctx context.Context, issueKey, targetQueue string, notify, notifyAuthor, moveAllFields, initialStatus bool, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("queue", targetQueue)
	q.Set("notify", strconv.FormatBool(notify))
	q.Set("notifyAuthor", strconv.FormatBool(notifyAuthor))
	q.Set("moveAllFields", strconv.FormatBool(moveAllFields))
	q.Set("initialStatus", strconv.FormatBool(initialStatus))
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/_move", q, fields)
}

// SuggestIssues returns issue suggestions matching input text, optionally scoped to a queue.
func (c *Client) SuggestIssues(ctx context.Context, input, queue string) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("input", input)
	if queue != "" {
		q.Set("queue", queue)
	}
	return c.do(ctx, http.MethodGet, "/v3/issues/_suggest", q, nil)
}

// ReleaseSearchScroll releases server-side resources held by a scrolled search-issues request.
func (c *Client) ReleaseSearchScroll(ctx context.Context, scrollID, scrollToken string) (json.RawMessage, error) {
	body := map[string]any{"scrollId": scrollID, "scrollToken": scrollToken}
	return c.do(ctx, http.MethodPost, "/v3/system/search/scroll/_clear", nil, body)
}

// AddChecklistItem adds an item to an issue's checklist, creating the checklist if needed.
func (c *Client) AddChecklistItem(ctx context.Context, issueKey string, item map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/checklistItems", nil, item)
}

// ListChecklistItems lists the items of an issue's checklist.
func (c *Client) ListChecklistItems(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/checklistItems", nil, nil)
}

// EditChecklistItems bulk-edits items of an issue's checklist; each item must include its id.
func (c *Client) EditChecklistItems(ctx context.Context, issueKey string, items []map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPatch, "/v3/issues/"+url.PathEscape(issueKey)+"/checklistItems", nil, items)
}

// DeleteChecklist deletes an issue's entire checklist.
func (c *Client) DeleteChecklist(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/issues/"+url.PathEscape(issueKey)+"/checklistItems", nil, nil)
}

// DeleteChecklistItem deletes a single item from an issue's checklist.
func (c *Client) DeleteChecklistItem(ctx context.Context, issueKey, itemID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/issues/"+url.PathEscape(issueKey)+"/checklistItems/"+url.PathEscape(itemID), nil, nil)
}

// GetComment retrieves a single issue comment.
func (c *Client) GetComment(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/comments/"+url.PathEscape(commentID), nil, nil)
}

// EditComment changes the text (and optionally summonees) of an existing comment.
func (c *Client) EditComment(ctx context.Context, issueKey, commentID, text string, summonees []string) (json.RawMessage, error) {
	body := map[string]any{"text": text}
	if len(summonees) > 0 {
		body["summonees"] = summonees
	}
	return c.do(ctx, http.MethodPatch, "/v3/issues/"+url.PathEscape(issueKey)+"/comments/"+url.PathEscape(commentID), nil, body)
}

// DeleteComment deletes an issue comment.
func (c *Client) DeleteComment(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/issues/"+url.PathEscape(issueKey)+"/comments/"+url.PathEscape(commentID), nil, nil)
}

// AddCommentReaction adds a reaction (e.g. "LIKE", "HEART") to a comment.
func (c *Client) AddCommentReaction(ctx context.Context, issueKey, commentID, reaction string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/comments/"+url.PathEscape(commentID)+"/reactions/"+url.PathEscape(reaction), nil, nil)
}

// ListIssueLinksPaginate lists issue links a page at a time, optionally filtered by link type.
func (c *Client) ListIssueLinksPaginate(ctx context.Context, issueKey string, perPage, page int, linkTypes []string) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	body := map[string]any{}
	if len(linkTypes) > 0 {
		body["linkTypes"] = linkTypes
	}
	return c.do(ctx, http.MethodPost, "/v3/issues/"+url.PathEscape(issueKey)+"/links/_list", q, body)
}

// DeleteIssueLink removes a link between two issues.
func (c *Client) DeleteIssueLink(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/issues/"+url.PathEscape(issueKey)+"/links/"+url.PathEscape(linkID), nil, nil)
}

// ListAttachments lists the files attached to an issue.
func (c *Client) ListAttachments(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/attachments", nil, nil)
}

// GetAttachmentInfo retrieves metadata for a single attached file.
func (c *Client) GetAttachmentInfo(ctx context.Context, issueKey, attachmentID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/issues/"+url.PathEscape(issueKey)+"/attachments/"+url.PathEscape(attachmentID), nil, nil)
}
