package model

import . "shared/option"

type Note int
const (
	C Note = iota
	CSharp
	DFlat
	D
	DSharp
	EFlat
	E
	F
	FSharp
	GFlat
	G
	GSharp
	AFlat
	A
	ASharp
	BFlat
	B
)

// no more modes for now
type KeyMode int
const (
	Major KeyMode = iota
	Minor
)

type WorkTitle struct {
	Kind string `json:"kind"`
	Number Option[int] `json:"number"` // it may be the only piece of a gender made by the composer
	Nickname Option[string] `json:"nickname"`
}

type Key struct {
	Note Note `json:"note"`
    Mode KeyMode `json:"mode"`
}

type Composer struct {
    Id int `json:"id"`
    Name string `json:"name"`
    BirthYear int `json:"birth_year"`
    DeathYear Option[int] `json:"death_year"` // the composer may be alive not
    PhotoPath string `json:"photo_path"` // relative to public/images/composers
}

// e.g.: Op. 27, No. 2 (Chopin's Nocturne in D-flat Major)
// the app assumes every work is catalogued, so this is mandatory
type Catalog struct {
	Prefix string `json:"prefix"` // Op., BWV, D., K...
	Number Option[string] `json:"number"` // it may be posthumous or it just doesn't have a number. It's a string because it may contain letters.
    Subnumber Option[string] `json:"subnumber"` // it may not have a subnumber. It may also contain letters.
}

type Movement struct {
	Nickname Option[string] `json:"nickname"` // e.g. "Ode to Joy"
	Form Option[string] `json:"form"` // e.g. Scherzo, Finale, Alla breve
	Lyrics Option[[]string] `json:"lyrics"`
	TempoMarkings []TempoMarking `json:"tempo_markings"`
	Sheet SheetMusic `json:"sheet"`
}

// TODO: add BPM
type TempoMarking struct {
	Form Option[string] `json:"form"` // e.g. Lassan, Friska, Fugue
	Name string `json:"name"`
}

type CompositionYear struct {
	StartYear int `json:"start_year"`
	EndYear Option[int] `json:"end_year"` // if the work has been finished in the same year of start
}

type SheetMusic struct {
	Path string `json:"path"` // relative to /public/sheets
}

type Work struct {
	Id int `json:"id"`
	Title WorkTitle `json:"title"`
	Key Key `json:"key"`
	Composer Composer `json:"composer"`
	Catalog Catalog `json:"catalog"`
	Movements []Movement `json:"movements"`
	Year CompositionYear `json:"year"`
}
