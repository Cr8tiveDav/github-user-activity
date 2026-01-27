package main

import (
	"encoding/json"
	"time"
)

// import "fmt"

type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Actor     Actor           `json:"actor"`
	Repo      Repo            `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	Public    bool            `json:"public"`
	CreatedAt time.Time       `json:"created_at"`
}

type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Events []Event

type PushEventPayload struct {
	Ref string `json:"ref"`
}

type CreateEventPayload struct {
	Ref     string `json:"ref,omitempty"`
	RefType string `json:"ref_type"`
}

type WatchEventPayload struct {
	Action string `json:"action"`
}

type DeleteEventPayload struct {
	CreateEventPayload
}

type ForkEventPayload struct {
	Forkee struct {
		FullName string `json:"full_name"`
	} `json:"forkee"`
}

type IssuesEventPayload struct {
	Action string `json:"action"`
	Issue  struct {
		Number int    `json:"number"`
		Title  string `json:"title"`
	} `json:"issue"`
}

type IssueCommentEventPayload struct {
	Issue struct {
		Number int `json:"number"`
	} `json:"issue"`
}

type PullRequest struct {
	Title  string `json:"title"`
	Number int    `json:"number,omitempty"`
}

type PullRequestEventPayload struct {
	Action      string `json:"action"`
	Number      int    `json:"number"`
	PullRequest `json:"pull_request"`
}

type PullRequestReviewEventPayload struct {
	Review struct {
		State string `json:"state"`
	} `json:"review"`
	PullRequest `json:"pull_request"`
}

type ReleaseEventPayload struct {
	Action  string `json:"action"`
	Release struct {
		TagName string `json:"tag_name"`
		Name    string `json:"name"`
	} `json:"release"`
}
