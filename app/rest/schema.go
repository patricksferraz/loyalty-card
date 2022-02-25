package rest

import (
	"time"
)

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HTTPResponse struct {
	Msg string `json:"msg,omitempty" example:"any message"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type CreateGuestRequest struct {
	Name string `json:"name"`
	Doc  string `json:"doc,omitempty"`
}

type Guest struct {
	Base               `json:",inline"`
	CreateGuestRequest `json:",inline"`
}

type CreateScoreRequest struct {
	Date        time.Time `json:"date" time_format:"RFC3339"`
	Description string    `json:"description,omitempty"`
	GuestID     string    `json:"guest_id"`
}

type Score struct {
	Base               `json:",inline"`
	CreateGuestRequest `json:",inline"`
	WasUsed            bool      `json:"was_used"`
	UsedIn             time.Time `json:"used_in,omitempty"`
	Tags               []Tag     `json:"tags,omitempty"`
}

type CreateTagRequest struct {
	Name string `json:"name"`
}

type Tag struct {
	Base             `json:",inline"`
	CreateTagRequest `json:",inline"`
}

type AddTagRequest struct {
	TagID string `json:"tag_id"`
}

type SearchTagsRequest struct {
	Name string `query:"name"`
}
