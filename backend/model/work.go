package model

type Note int
const (
	C Note = iota
	CSharp
	DFlat
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
	Number int
	Nickname *string // optional
}

type Key struct {
	CentralNote Note
    Mode KeyMode
}

type Composer struct {
    Name string
	BirthDate int
	DeathDate *int // optional, the composer may be alive now
}

type TempoMarking struct {
	Name string
	Bpm *int // optional, since many markings don’t have a strict BPM
}

// e.g.: Op. 27, No. 2 (Chopin's Nocturne in D-flat Major)
type Catalog struct {
	Name string // Op, D, K..
	Number int
    InsideNum *int // optional, it may not have a number.
}

type Movement struct {
	Kind *string // optional, e.g. Scherzo, Finale, Alla breve
	TempoMarkings []TempoMarking
	Nickname *string // optional
	Order int
}

type CompositionYear struct {
	StartYear int
	EndYear *int // optional, if the work has been finished in the same year of start
}

type Work struct {
	Id int
	Title WorkTitle
	Key Key
	Composer Composer
	Catalog Catalog
	Movements []Movement
	Year CompositionYear
}
