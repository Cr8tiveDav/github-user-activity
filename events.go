package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func printEvents(events Events, username string) {
	if len(events) == 0 {
		fmt.Printf("No recent public activities found for github user: %s\n", username)
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\nRecent GitHub event for %s:\n\n", username)
	// Print Header
	fmt.Fprintln(w, "Date\tEvent\tActivity")
	// Print Separator
	fmt.Fprintln(w, "---\t---\t---")

	for _, event := range events {
		date := event.CreatedAt.Format("2006-01-02 15:04")
		switch event.Type {
		case "CreateEvent":
			var payload CreateEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\tCreated %s %s in **%s**\n",
				date,
				event.Type,
				payload.RefType,
				payload.Ref,
				event.Repo.Name)
		case "PushEvent":
			var payload PushEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\tPushed to %s in **%s**\n",
				date,
				event.Type,
				payload.Ref,
				event.Repo.Name)
		case "WatchEvent":
			fmt.Fprintf(w, "%s\t%s\tStarred **%s**\n",
				date,
				event.Type,
				event.Repo.Name)
		case "ForkEvent":
			var payload ForkEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\t Forked **%s** to **%s**\n",
				date,
				event.Type,
				event.Repo.Name,
				payload.Forkee.FullName)
		case "DeleteEvent":
			var payload DeleteEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\tDelete %s **%s** in **%s**\n",
				date,
				event.Type,
				payload.RefType,
				payload.Ref,
				event.Repo.Name)
		case "IssuesEvent":
			var payload IssuesEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\t%s issue **#%d** %s in **%s**\n",
				date,
				event.Type,
				strings.ToUpper(payload.Action[:1])+payload.Action[1:],
				payload.Issue.Number,
				payload.Issue.Title,
				event.Repo.Name)
		case "IssueCommentEvent":
			var payload IssueCommentEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\tCommented on issue **#%d** in **%s**\n",
				date,
				event.Type,
				payload.Issue.Number,
				event.Repo.Name)
		case "Pull`RequestEvent":
			var payload PullRequestEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\t %s pull request **#%d** %s in **%s**\n",
				date,
				event.Type,
				strings.ToUpper(payload.Action[:1])+payload.Action[1:],
				payload.Number,
				payload.PullRequest.Title,
				event.Repo.Name)
		case "PullRequestReviewEvent":
			var payload PullRequestReviewEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\t%s review on pull request **#%d** in **%s**\n",
				date,
				event.Type,
				strings.ToUpper(payload.Review.State)[:1]+payload.Review.State[1:],
				payload.PullRequest.Number,
				event.Repo.Name)
		case "ReleaseEvent":
			var payload ReleaseEventPayload
			json.Unmarshal(event.Payload, &payload)
			fmt.Fprintf(w, "%s\t%s\t%s release **%s** - **%s** in **%s**\n",
				date,
				event.Type,
				strings.ToUpper(payload.Action[:1])+payload.Action[1:],
				payload.Release.TagName,
				payload.Release.Name,
				event.Repo.Name)
		default:
			fmt.Fprintf(w, "%s\t%s\t%s in **%s**\n",
				date,
				event.Type,
				event.Type,
				event.Repo.Name)
		}
	}
}
