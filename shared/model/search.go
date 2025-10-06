package model

import (
	"encoding/json"
)

type SearchResult interface {
	MarshalJSON() ([]byte, error)
}

// ---

func (w Work) MarshalJSON() ([]byte, error) {
	type Content Work
	return json.Marshal(&struct {
		Type string `json:"type"`
		Content
	}{
		Type: "work",
		Content: Content(w),
	})
}

func (c Composer) MarshalJSON() ([]byte, error) {
	type Content Composer
	return json.Marshal(&struct {
		Type string `json:"type"`
		Content
	}{
		Type: "composer",
		Content: Content(c),
	})
}

func (p Performer) MarshalJSON() ([]byte, error) {
	type Content Performer
	return json.Marshal(&struct {
		Type string `json:"type"`
		Content
	}{
		Type: "performer",
		Content: Content(p),
	})
}
