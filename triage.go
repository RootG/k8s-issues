package main

import (
	"context"
	"net/http"
)

// An issue needs triage if:
//	* Is labeled sig/network (they've opted in to this)
//	* Does not already have label triage/unresolved
//  * TODO: Less than 24h old?
//	* Does not have "/triage resolved" comment
func issueNeedsTriage(issue *Issue) bool {
	// Only sig-network has opted in.
	if !issue.hasLabel("sig/network") {
		return false
	}

	// Don't double-comment.
	if issue.hasLabel("triage/unresolved") {
		return false
	}

	// Don't relabel resolved issues.
	if issue.hasCommentWithCommand("/triage", "resolved") {
		return false
	}

	return true
}

func triageLabel(ctx context.Context, httpClient *http.Client, issue *Issue) {
	if issueNeedsTriage(issue) {
		addComment(ctx, httpClient,issue.Id,"/triage unresolved")
	}
}