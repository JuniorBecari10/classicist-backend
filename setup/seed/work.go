package seed

import (
	"shared/model"
	. "shared/option"
)

// Do not fill the ID.
var Composers = []model.Composer{
	{
		Name: "Johann Sebastian Bach",
		BirthYear: 1685,
		DeathYear: Some(1750),
		PhotoPath: "bach.jpg",
	},
	{
		Name: "Antonio Vivaldi",
		BirthYear: 1678,
		DeathYear: Some(1741),
		PhotoPath: "vivaldi.jpg",
	},
	{
		Name: "Wolfgang Amadeus Mozart",
		BirthYear: 1756,
		DeathYear: Some(1791),
		PhotoPath: "mozart.jpg",
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
}

// Do not fill the ID.
var Works = []model.Work{
	{
		Title: model.WorkTitle{
			Kind: "Prelude",
			Number: Some(1),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.C,
			Mode: model.Major,
		},

		Composer: Composers[0], // Bach

		Catalog: model.Catalog{
			Prefix: "BWV",
			Number: Some("846"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "prelude.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "[no tempo marking]",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1722,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Toccata and Fugue",
			Number: None[int](),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.D,
			Mode: model.Minor,
		},

		Composer: Composers[0], // Bach

		Catalog: model.Catalog{
			Prefix: "BWV",
			Number: Some("565"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "toccata.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "[no tempo marking]",
					},
					{
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1704,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Violin Concerto",
			Number: Some(1),
			Nickname: Some("Spring"),
		},

		Key: model.Key{
			Note: model.E,
			Mode: model.Major,
		},

		Composer: Composers[1], // Vivaldi

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("8"),
			Subnumber: Some("1"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "spring-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "spring-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Largo",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "spring-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1723,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Violin Concerto",
			Number: Some(2),
			Nickname: Some("Summer"),
		},

		Key: model.Key{
			Note: model.G,
			Mode: model.Minor,
		},

		Composer: Composers[1], // Vivaldi

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("8"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "summer-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro non molto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "summer-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio e piano",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "summer-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Presto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1723,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Violin Concerto",
			Number: Some(3),
			Nickname: Some("Autumn"),
		},

		Key: model.Key{
			Note: model.F,
			Mode: model.Major,
		},

		Composer: Composers[1], // Vivaldi

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("8"),
			Subnumber: Some("3"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "autumn-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "autumn-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio molto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "autumn-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1704,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Violin Concerto",
			Number: Some(4),
			Nickname: Some("Winter"),
		},

		Key: model.Key{
			Note: model.F,
			Mode: model.Minor,
		},

		Composer: Composers[1], // Vivaldi

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("8"),
			Subnumber: Some("4"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "winter-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro non molto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "winter-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Largo",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "winter-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1704,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Serenade",
			Number: Some(13),
			Nickname: Some("Eine kleine Nachtmusik"),
		},

		Key: model.Key{
			Note: model.G,
			Mode: model.Major,
		},

		Composer: Composers[2], // Mozart

		Catalog: model.Catalog{
			Prefix: "K.",
			Number: Some("525"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "eine-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Romanze"),
				
				Sheet: model.SheetMusic{
					Path: "eine-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Menuetto"),
				
				Sheet: model.SheetMusic{
					Path: "eine-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Rondo"),
				
				Sheet: model.SheetMusic{
					Path: "eine-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1787,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Symphony",
			Number: Some(40),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.G,
			Mode: model.Minor,
		},

		Composer: Composers[2], // Mozart

		Catalog: model.Catalog{
			Prefix: "K.",
			Number: Some("550"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "symp40-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Molto allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "symp40-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Menuetto"),
				
				Sheet: model.SheetMusic{
					Path: "symp40-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				
				Sheet: model.SheetMusic{
					Path: "symp40-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro assai",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1787,
			EndYear: None[int](),
		},
	},
	{
		Title: model.WorkTitle{
			Kind: "Piano Sonata",
			Number: Some(14),
			Nickname: Some("Moonlight"),
		},

		Key: model.Key{
			Note: model.CSharp,
			Mode: model.Minor,
		},

		Composer: Composers[3], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("27"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "moonlight-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio sostenuto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "moonlight-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "moonlight-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Presto agitato",
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

		Composer: Composers[3], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("67"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "symp5-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro con brio",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "symp5-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante con moto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Scherzo"),

				Sheet: model.SheetMusic{
					Path: "symp5-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Finale"),

				Sheet: model.SheetMusic{
					Path: "symp5-4.mxl",
				},

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

		Composer: Composers[3], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("92"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "symp7-1.mxl",
				},

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

				Sheet: model.SheetMusic{
					Path: "symp7-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "symp7-3.mxl",
				},

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

				Sheet: model.SheetMusic{
					Path: "symp7-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro con brio",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1811,
			EndYear: Some(1812),
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

		Composer: Composers[3], // Beethoven

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("125"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "symp9-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Allegro ma non troppo, un poco maestoso",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "symp9-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Molto vivace",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "symp9-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Adagio molto e cantabile",
					},
				},
			},
			{
				Nickname: Some("Ode to Joy"),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "symp9-4.mxl",
				},

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
			Kind: "Bagatelle",
			Number: Some(25),
			Nickname: Some("Für Elise"),
		},

		Key: model.Key{
			Note: model.A,
			Mode: model.Minor,
		},

		Composer: Composers[3], // Beethoven

		Catalog: model.Catalog{
			Prefix: "WoO",
			Number: Some("59"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "furelise.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Poco moto",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1810,
			EndYear: None[int](),
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

		Composer: Composers[4], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("9"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "nocturne2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Andante",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1832,
			EndYear: None[int](),
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

		Composer: Composers[4], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op. posth.",
			Number: None[string](),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "nocturne20.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lento con gran espressione",
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
			Kind: "Piano Sonata",
			Number: Some(2),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.BFlat,
			Mode: model.Minor,
		},

		Composer: Composers[4], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("35"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "chopin-sonata2-1.mxl",
				},

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

				Sheet: model.SheetMusic{
					Path: "chopin-sonata2-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Scherzo",
					},
				},
			},
			{
				Nickname: Some("Funeral March"),
				Form: Some("Marche funèbre"),

				Sheet: model.SheetMusic{
					Path: "chopin-sonata2-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lento",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Finale"),

				Sheet: model.SheetMusic{
					Path: "chopin-sonata2-4.mxl",
				},

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
			Kind: "Ballade",
			Number: Some(1),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.G,
			Mode: model.Minor,
		},

		Composer: Composers[4], // Chopin

		Catalog: model.Catalog{
			Prefix: "Op.",
			Number: Some("23"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "ballade1.mxl",
				},

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
			Kind: "Liebestraum",
			Number: Some(3),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.AFlat,
			Mode: model.Major,
		},

		Composer: Composers[5], // Liszt

		Catalog: model.Catalog{
			Prefix: "S.",
			Number: Some("541"),
			Subnumber: Some("3"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: Some("Notturno"),

				Sheet: model.SheetMusic{
					Path: "liebestraum3.mxl",
				},

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
			Kind: "Hungarian Rhapsody",
			Number: Some(2),
			Nickname: None[string](),
		},

		Key: model.Key{
			Note: model.CSharp,
			Mode: model.Minor,
		},

		Composer: Composers[5], // Liszt

		Catalog: model.Catalog{
			Prefix: "S.",
			Number: Some("244"),
			Subnumber: Some("2"),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),

				Sheet: model.SheetMusic{
					Path: "hungrhap2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Name: "Lento a capriccio",
					},
					{
						Name: "Lassan: Andante mesto",
					},
					{
						Name: "Friska: Vivace",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1847,
			EndYear: None[int](),
		},
	},
}
