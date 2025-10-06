package model

type SearchItem interface {
	search()
}

type SearchResult struct {
	Type string `json:"type"`
	Data SearchItem `json:"data"`
}

// ---

func NewWorkSR(w Work) SearchResult {
	return SearchResult{
		Type: "work",
		Data: w,
	}
}

func NewComposerSR(c Composer) SearchResult {
	return SearchResult{
		Type: "composer",
		Data: c,
	}
}

func NewPerformerSR(p Performer) SearchResult {
	return SearchResult{
		Type: "performer",
		Data: p,
	}
}

// ---

func (x Work) search() { }
func (x Composer) search() { }
func (x Performer) search() { }
