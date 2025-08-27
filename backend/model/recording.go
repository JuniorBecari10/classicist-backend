package model

type Performer struct {
	Name string
	Role string // e.g. Conductor, Pianist, Violinist, Orchestra
}

type RecordedMovement struct {
	MovementId int
	RecordingId int
	File AudioFile
}

type Recording struct {
	Id int
	WorkId int // what is the work this recording is from.

	Performers []Performer
	Year int
	Movements []RecordedMovement
}

type AudioFile struct {
	Id int
	Path string
	Duration int // in seconds
}
