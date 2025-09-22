package seed

import (
	"shared/model"
	. "shared/option"
)

var Composers = []model.Composer{
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

var Works = []model.Work{
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

		Composer: Composers[0], // Bach

		Catalog: model.Catalog{
			Prefix: "BWV",
            Number: Some("1052"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},

				Sheet: model.SheetMusic{
					Path: "bhc1-1.mxl",
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio",
					},
				},

				Sheet: model.SheetMusic{
					Path: "bhc1-2.mxl",
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},

				Sheet: model.SheetMusic{
					Path: "bhc1-3.mxl",
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

		Composer: Composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("73"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio un poco mosso",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "bpc5.mxl",
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

		Composer: Composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("31"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "bps17.mxl",
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

		Composer: Composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("67"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro con brio",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante con moto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Scherzo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
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
		
		Sheet: model.SheetMusic{
			Path: "bs5.mxl",
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

		Composer: Composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("92"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
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
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "bs7.mxl",
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

		Composer: Composers[1], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("125"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro ma non troppo, un poco maestoso",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Molto vivace",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio molto e cantabile",
					},
				},
			},
			{
				Nickname: Some("Ode to Joy"),
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

		Sheet: model.SheetMusic{
			Path: "bs9.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("35"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Scherzo",
					},
				},
			},
			{
				Nickname: Some("Funeral March"),
				Form: Some("Marche funèbre"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lento",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cps2.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("58"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro maestoso",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Molto vivace",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Largo",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cps3.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("23"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cb1.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("52"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cb4.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("9"),
			Subnumber: Some("1"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cn1.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("9"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cn2.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("15"),
			Subnumber: Some("1"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cn4.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("27"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cn8.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op. posth.",
			Number: None[string](),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cn20.mxl",
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

		Composer: Composers[2], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("11"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro maestoso",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Larghetto",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "cpc1.mxl",
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

		Composer: Composers[3], // Liszt

		Catalog: model.Catalog{
			Prefix: "S.",
			Number: Some("541"),
			Subnumber: Some("3"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "ll3.mxl",
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

		Composer: Composers[6], // Mozart

		Catalog: model.Catalog{
			Prefix: "K.",
			Number: Some("466"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Romanze",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "mpc20.mxl",
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

		Composer: Composers[5], // Mendelssohn

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("52"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: Some("Sinfonia"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Maestoso con moto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto un poco agitato",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio religioso",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Chorale"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "O Herr, lehre mich",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Chorus"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Nun danket alle Gott",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Aria"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Ihr Sorgen, flieht",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Chorus"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Wer bis an das Ende beharrt",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Aria"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Ihr Völker des Himmels",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Chorus"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Halleluja",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Aria"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Preiset den Herrn",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Chorus"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lobet den Herrn",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "mds2.mxl",
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

		Composer: Composers[5], // Mendelssohn

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("64"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro molto appassionato",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "mvc.mxl",
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

		Composer: Composers[4], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("30"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro maestoso",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante moderato",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "In ruhig fließender Bewegung",
					},
					
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Urllicht"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Sehr feierlich, aber schlicht",
					},
					
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "mas2.mxl",
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

		Composer: Composers[4], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("31"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Kräftig. Entschieden",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Tempo di Menuetto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Comodo. Scherzando",
					},
					
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Sehr langsam",
					},
					
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Misterioso",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lustig im Tempo und keck im Ausdruck",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "mas3.mxl",
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

		Composer: Composers[4], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("44"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: Some("Traeurmarsch"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "In gemessenem Schritt. Streng. Wie ein Kondukt",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Stürmisch bewegt. Mit grösster Vehemenz",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Scherzo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Kräftig, Nicht zu schnell",
					},
					
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagietto. Sehr langsam",
					},
					
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "mas5.mxl",
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

		Composer: Composers[4], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("48"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: Some("Part I"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Veni, Creator Spiritus",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "mas8.mxl",
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

		Composer: Composers[4], // Mahler

		Catalog: model.Catalog{
			Prefix: "GMW",
			Number: Some("F58"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Scherzo",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Purgatorio"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto moderato",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Scherzo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro pesante",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "mas10.mxl",
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

		Composer: Composers[7], // Rachmaninoff

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("18"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Moderato",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio sostenuto",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "rpc2.mxl",
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

		Composer: Composers[7], // Rachmaninoff

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("30"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro ma non tanto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Intermezzo"),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "rpc3.mxl",
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

		Composer: Composers[8], // Schubert

		Catalog: model.Catalog{
			Prefix: "D.",
			Number: Some("759"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro moderato",
					},
				},
			},
			{
				Nickname: None[string](),
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

		Sheet: model.SheetMusic{
			Path: "ss8.mxl",
		},
	},
}
