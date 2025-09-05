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
	Kind string
	Number Option[int] // it may be the only piece of a gender made by the composer
	Nickname Option[string]
}

type Key struct {
	Note Note
    Mode KeyMode
}

type Composer struct {
    Name string
	BirthYear int
	DeathYear Option[int] // the composer may be alive now
	PhotoPath string // relative to /public/images/composers
}

// e.g.: Op. 27, No. 2 (Chopin's Nocturne in D-flat Major)
// the app assumes every work is catalogued, so this is mandatory
type Catalog struct {
	Prefix string // Op., BWV, D., K...
	Number Option[string] // it may be posthumous or it just doesn't have a number. It's a string because it may contain letters.
    Subnumber Option[string] // it may not have a subnumber. It may also contain letters.
}

type Movement struct {
	Kind Option[string] // e.g. Scherzo, Finale, Alla breve
	TempoMarkings []TempoMarking
	Order int
}

// TODO: add BPM
type TempoMarking struct {
	Name string
}

type CompositionYear struct {
	StartYear int
	EndYear Option[int] // if the work has been finished in the same year of start
}

type Work struct {
	Title WorkTitle
	Key Key
	Composer Composer
	Catalog Catalog
	Movements []Movement
	Year CompositionYear
}
