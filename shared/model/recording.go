package model

type AudioFile struct {
	Path string // relative to /public/audio
	Duration int // in seconds
}

type Performer struct {
	Id int
	Name string
}

// two different structs are necessary because a single performer
// can have multiple roles, depending on the recording.
type RecordingPerformer struct {
	Id int
	Performer Performer
	Role string // e.g. Conductor, Pianist, Violinist, Orchestra
}

type RecordedMovement struct {
	Id int
	MovementIndex int // to avoid data duplication, it is already in Work, inside Recording. Use the index to get it.
	MovementId int
	AudioFile AudioFile
}

type Recording struct {
	Id int
	WorkId int
	Year int
	PhotoPath string // relative to /public/images/recordings
	Performers []RecordingPerformer
	Movements []RecordedMovement
}
