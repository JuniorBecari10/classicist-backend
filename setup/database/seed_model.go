package database

import (
	"shared/model"
	. "shared/option"
)

var composers = []model.Composer{
	{
		Name: "Johann Sebastian Bach",
		BirthYear: 1685,
		DeathYear: Some(1750),
		PhotoPath: "bach.jpg",
	},
	{
		Name: "Ludwig van Beethoven",
		BirthYear: 1770,
		DeathYear: Some(1827),
		PhotoPath: "beethoven.jpg",
	},
	{
		Name: "Frédéric Chopin",
		BirthYear: 1810,
		DeathYear: Some(1849),
		PhotoPath: "chopin.jpg",
	},
	{
		Name: "Franz Liszt",
		BirthYear: 1811,
		DeathYear: Some(1886),
		PhotoPath: "liszt.jpg",
	},
	{
		Name: "Gustav Mahler",
		BirthYear: 1860,
		DeathYear: Some(1911),
		PhotoPath: "mahler.jpg",
	},
	{
		Name: "Felix Mendelssohn",
		BirthYear: 1809,
		DeathYear: Some(1847),
		PhotoPath: "mendelssohn.jpg",
	},
	{
		Name: "Wolfgang Amadeus Mozart",
		BirthYear: 1756,
		DeathYear: Some(1791),
		PhotoPath: "mozart.jpg",
	},
	{
		Name: "Sergei Rachmaninoff",
		BirthYear: 1873,
		DeathYear: Some(1943),
		PhotoPath: "rachmaninoff.jpg",
	},
	{
		Name: "Franz Schubert",
		BirthYear: 1797,
		DeathYear: Some(1828),
		PhotoPath: "schubert.jpg",
	},
}

var works = []model.Work{
	{
		Title: model.WorkTitle{
			Kind: "Harpsichord Concerto",
			Number: Some(1),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.D,
			Mode: model.Minor,
		},

		Composer: composers[0], // Bach

		Catalog: model.Catalog{
			Prefix: "BWV",
            Number: Some("1052"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1730,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Concerto",
			Number: Some(5),
			Nickname: Some("Emperor"),
		},

		Key: model.Key{
			Note: model.EFlat,
			Mode: model.Major,
		},

		Composer: composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("73"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio un poco mosso",
					},
				},
			},
			{
				Form: Some("Rondo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1809,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Sonata",
			Number: Some(17),
			Nickname: Some("Tempest"),
		},

		Key: model.Key{
			Note: model.D,
			Mode: model.Minor,
		},

		Composer: composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("31"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Largo",
					},
					{
						Name: "Allegro",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1801,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(5),
			Nickname: Some("Fate"),
		},

		Key: model.Key{
			Note: model.C,
			Mode: model.Minor,
		},

		Composer: composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("67"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro con brio",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante con moto",
					},
				},
			},
			{
				Form: Some("Scherzo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Form: Some("Finale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1804,
			EndYear: Some(1808),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(7),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.A,
			Mode: model.Major,
		},

		Composer: composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("92"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Poco sostenuto",
					},
					{
						Name: "Vivace",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Presto",
					},
					{
						Name: "Assai meno presto",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro con brio",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1811,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(9),
			Nickname: Some("Choral"),
		},

		Key: model.Key{
			Note: model.D,
			Mode: model.Minor,
		},

		Composer: composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("125"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro ma non troppo, un poco maestoso",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Molto vivace",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio molto e cantabile",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Presto",
					},
					{
						Name: "Allegro assai",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1822,
			EndYear: Some(1824),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Sonata",
			Number: Some(2),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.BFlat,
			Mode: model.Minor,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("35"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Grave",
					},
					{
						Name: "Doppio movimento",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Scherzo",
					},
				},
			},
			{
				Form: Some("Marche funèbre"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lento",
					},
				},
			},
			{
				Form: Some("Finale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Presto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1839,
			EndYear: Some(1840),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Sonata",
			Number: Some(3),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.B,
			Mode: model.Minor,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("58"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro maestoso",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Molto vivace",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Largo",
					},
				},
			},
			{
				Form: Some("Finale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Presto non tanto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1844,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Ballade",
			Number: Some(1),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.G,
			Mode: model.Minor,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("23"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Largo",
					},
					{
						Name: "Moderato",
					},
					{
						Name: "Presto con fuoco",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1831,
			EndYear: Some(1835),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Ballade",
			Number: Some(4),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.F,
			Mode: model.Minor,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("52"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Largo",
					},
					{
						Name: "Moderato",
					},
					{
						Name: "Presto con fuoco",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1842,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Nocturne",
			Number: Some(1),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.BFlat,
			Mode: model.Minor,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("9"),
			Subnumber: Some("1"),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Larghetto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1830,
			EndYear: Some(1832),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Nocturne",
			Number: Some(2),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.EFlat,
			Mode: model.Major,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("9"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1830,
			EndYear: Some(1832),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Nocturne",
			Number: Some(4),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.F,
			Mode: model.Major,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("15"),
			Subnumber: Some("1"),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante cantabile",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1830,
			EndYear: Some(1833),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Nocturne",
			Number: Some(8),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.DFlat,
			Mode: model.Major,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("27"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lento sostenuto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1835,
			EndYear: Some(1836),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Nocturne",
			Number: Some(20),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.CSharp,
			Mode: model.Minor,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op. posth.",
			Number: None[string](),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lento sostenuto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1830,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Concerto",
			Number: Some(1),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.E,
			Mode: model.Minor,
		},

		Composer: composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("11"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro maestoso",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Larghetto",
					},
				},
			},
			{
				Form: Some("Rondo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Vivace",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1830,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Liebestraum",
			Number: Some(3),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.AFlat,
			Mode: model.Major,
		},

		Composer: composers[3], // Liszt

		Catalog: model.Catalog{
			Prefix: "S.",
			Number: Some("541"),
			Subnumber: Some("3"),
		},

		Movements: []model.Movement{
			{
				Form: Some("Notturno"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Poco allegro, con affetto",
					},
					{
						Name: "Tempo I",
					},
				},
			},
			
		},

		Year: model.CompositionYear{
			StartYear: 1850,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Concerto",
			Number: Some(20),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.D,
			Mode: model.Minor,
		},

		Composer: composers[6], // Mozart

		Catalog: model.Catalog{
			Prefix: "K.",
			Number: Some("466"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Romanze",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro assai",
					},
				},
			},
			
		},

		Year: model.CompositionYear{
			StartYear: 1850,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(2),
			Nickname: Some("Lobgesang"),
		},

		Key: model.Key{
			Note: model.BFlat,
			Mode: model.Major,
		},

		Composer: composers[5], // Mendelssohn

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("52"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: Some("Sinfonia"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Maestoso con moto",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto un poco agitato",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio religioso",
					},
				},
			},
			{
				Form: Some("Chorale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "O Herr, lehre mich",
					},
				},
			},
			{
				Form: Some("Chorus"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Nun danket alle Gott",
					},
				},
			},
			{
				Form: Some("Aria"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Ihr Sorgen, flieht",
					},
				},
			},
			{
				Form: Some("Chorus"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Wer bis an das Ende beharrt",
					},
				},
			},
			{
				Form: Some("Aria"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Ihr Völker des Himmels",
					},
				},
			},
			{
				Form: Some("Chorus"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Halleluja",
					},
				},
			},
			{
				Form: Some("Aria"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Preiset den Herrn",
					},
				},
			},
			{
				Form: Some("Chorus"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lobet den Herrn",
					},
				},
			},
			{
				Form: Some("Finale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Alles was Odem hat, lobe den Herrn",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1840,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Violin Concerto",
			Number: None[int](),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.E,
			Mode: model.Minor,
		},

		Composer: composers[5], // Mendelssohn

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("64"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro molto appassionato",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto non troppo",
					},
					{
						Name: "Allegro molto vivace",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1844,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(2),
			Nickname: Some("Resurrection"),
		},

		Key: model.Key{
			Note: model.C,
			Mode: model.Minor,
		},

		Composer: composers[7], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("30"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro maestoso",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante moderato",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "In ruhig fließender Bewegung",
					},
					
				},
			},
			{
				Form: Some("Urllicht"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Sehr feierlich, aber schlicht",
					},
					
				},
			},
			{
				Form: Some("Finale. Im Tempo des Scherzos"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Wild herausfahrend",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1888,
			EndYear: Some(1894),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(3),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.D,
			Mode: model.Minor,
		},

		Composer: composers[7], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("31"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Kräftig. Entschieden",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Tempo di Menuetto",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Comodo. Scherzando",
					},
					
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Sehr langsam",
					},
					
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Misterioso",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lustig im Tempo und keck im Ausdruck",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Langsam. Ruhevoll. Empfunden",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1893,
			EndYear: Some(1896),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(5),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.CSharp,
			Mode: model.Minor,
		},

		Composer: composers[7], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("44"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Traeurmarsch. In gemessenem Schritt. Streng. Wie ein Kondukt",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Stürmisch bewegt. Mit grösster Vehemenz",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Scherzo. Kräftig, Nicht zu schnell",
					},
					
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagietto. Sehr langsam",
					},
					
				},
			},
			{
				Form: Some("Rondo-Finale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro – Allegro giocoso. Frisch",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1901,
			EndYear: Some(1902),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(8),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.EFlat,
			Mode: model.Major,
		},

		Composer: composers[7], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("48"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: Some("Part I"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Veni, creator spiritus",
					},
				},
			},
			{
				Form: Some("Part II. Finale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Alles Vergängliche ist nur ein Gleichnis",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1906,
			EndYear: Some(1907),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(10),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.FSharp,
			Mode: model.Major,
		},

		Composer: composers[7], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("F58"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Scherzo",
					},
				},
			},
			{
				Form: Some("Purgatorio"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto moderato",
					},
				},
			},
			{
				Form: Some("Scherzo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro pesante",
					},
				},
			},
			{
				Form: Some("Finale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Langsam, schwer",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1906,
			EndYear: Some(1907),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Concerto",
			Number: Some(2),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.C,
			Mode: model.Minor,
		},

		Composer: composers[8], // Rachmaninoff

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("18"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Moderato",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio sostenuto",
					},
				},
			},
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro scherzando",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1900,
			EndYear: Some(1901),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Concerto",
			Number: Some(3),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.D,
			Mode: model.Minor,
		},

		Composer: composers[8], // Rachmaninoff

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("30"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro ma non tanto",
					},
				},
			},
			{
				Form: Some("Intermezzo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio",
					},
				},
			},
			{
				Form: Some("Finale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Alla breve",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1900,
			EndYear: Some(1901),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(8),
			Nickname: Some("Unfinished"),
		},

		Key: model.Key{
			Note: model.B,
			Mode: model.Minor,
		},

		Composer: composers[9], // Schubert

		Catalog: model.Catalog{
			Prefix: "D.",
			Number: Some("759"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro moderato",
					},
				},
			},
			{
				Form: Some("Intermezzo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante con moto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1822,
			EndYear: Some(1823),
		},
	},
}
