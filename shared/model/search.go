package model

type SearchItem interface {
	search()
}

type SearchResult struct {
	Type string `json:"type"`
	Value SearchItem `json:"value"`
}

// ---

func NewWorkSR(w Work) SearchResult {
	return SearchResult{
		Type: "work",
		Value: w,
	}
}

func NewComposerSR(c Composer) SearchResult {
	return SearchResult{
		Type: "composer",
		Value: c,
	}
}

func NewPerformerSR(p Performer) SearchResult {
	return SearchResult{
		Type: "performer",
		Value: p,
	}
}

// ---

func (x Work) search() {}
func (x Composer) search() {}
func (x Performer) search() {}
