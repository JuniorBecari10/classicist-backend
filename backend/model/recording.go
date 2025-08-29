package model

type AudioFile struct {
	Path string
	Duration int // in seconds
}

type Performer struct {
	Name string
}

// two different structs are necessary because a single performer
// can have multiple roles, depending on the recording.
type RecordingPerformer struct {
	Performer Performer
	Role string // e.g. Conductor, Pianist, Violinist, Orchestra
}

type RecordedMovement struct {
	Movement Movement // yes, a copy. no problem for now
	AudioFile AudioFile
}

type Recording struct {
	Work Work
	Year int
	PhotoPath string
	Performers []RecordingPerformer
	Movements []RecordedMovement
}
