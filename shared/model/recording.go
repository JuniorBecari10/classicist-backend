package model

type AudioFile struct {
	Path string `json:"path"` // relative to /public/audio
	Duration int `json:"duration"` // in seconds
}

type Performer struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

// two different structs are necessary because a single performer
// can have multiple roles, depending on the recording.
type RecordingPerformer struct {
	Id int `json:"id"`
	Performer Performer `json:"performer"`
	Role string `json:"role"` // e.g. Conductor, Pianist, Violinist, Orchestra
}

type RecordedMovement struct {
	Id int `json:"id"`
	MovementIndex int `json:"movement_index"` // use index to locate the movement inside Work
	MovementId int `json:"movement_id"`
	AudioFile AudioFile `json:"audio_file"`
}

type Recording struct {
	Id int `json:"id"`
	WorkId int `json:"work_id"`
	Year int `json:"year"`
	PhotoPath string `json:"photo_path"` // relative to /public/images/recordings
	Performers []RecordingPerformer `json:"performers"`
	Movements []RecordedMovement `json:"movements"`
}
