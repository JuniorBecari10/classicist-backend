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
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "prelude.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "toccata.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: Some("Toccata"),
						Name: "[no tempo marking]",
					},
					{
						Form: Some("Fugue"),
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
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "spring-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "spring-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Largo",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "spring-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "summer-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro non molto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "summer-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Adagio e piano",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "summer-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "autumn-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "autumn-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Adagio molto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "autumn-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "winter-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro non molto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "winter-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Largo",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "winter-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "eine-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Romanze"),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "eine-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Andante",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Menuetto"),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "eine-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Rondo"),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "eine-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "symp40-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Molto allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "symp40-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Andante",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Menuetto"),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "symp40-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),				Lyrics: None[[]string](),


				Sheet: model.SheetMusic{
					Path: "symp40-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
			Number: Some(16),
			Nickname: Some("Sonata facile"),
		},

		Key: model.Key{
			Note: model.C,
			Mode: model.Major,
		},

		Composer: Composers[2], // Mozart

		Catalog: model.Catalog{
			Prefix: "K.",
			Number: Some("545"),
			Subnumber: None[string](),
		},

		Movements: []model.Movement{
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "facile-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "facile-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Andante",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Rondo"),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "facile-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro",
					},
				},
			},
		},

		Year: model.CompositionYear{
			StartYear: 1788,
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "moonlight-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Adagio sostenuto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "moonlight-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "moonlight-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp5-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro con brio",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp5-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Andante con moto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Scherzo"),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp5-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Finale"),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp5-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp7-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Poco sostenuto",
					},
					{
						Form: None[string](),
						Name: "Vivace",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp7-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegretto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp7-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Presto",
					},
					{
						Form: None[string](),
						Name: "Assai meno presto",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp7-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp9-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Allegro ma non troppo, un poco maestoso",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp9-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Molto vivace",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "symp9-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Adagio molto e cantabile",
					},
				},
			},
			{
				Nickname: Some("Ode to Joy"),
				Form: None[string](),

				Lyrics: Some([]string{
					"Freude, schöner Götterfunken,",
					"Tochter aus Elysium,",
					"Wir betreten feuertrunken,",
					"Himmlische, dein Heiligtum.",
					"Deine Zauber binden wieder,",
					"Was die Mode streng geteilt,*",
					"Alle Menschen werden Brüder,*",
					"Wo dein sanfter Flügel weilt.",
					"Seid umschlungen Millionen!",
					"Diesen Kuß der ganzen Welt!",
					"Brüder – überm Sternenzelt",
					"Muß ein lieber Vater wohnen.",
					"Wem der große Wurf gelungen,",
					"Eines Freundes Freund zu sein,",
					"Wer ein holdes Weib errungen,",
					"Mische seinen Jubel ein!",
					"Ja – wer auch nur eine Seele",
					"Sein nennt auf dem Erdenrund!",
					"Und wer's nie gekonnt, der stehle",
					"Weinend sich aus diesem Bund!",
					"Was den großen Ring bewohnet",
					"Huldige der Sympathe!",
					"Zu den Sternen leitet sie,",
					"Wo der Unbekannte thronet.",
					"Freude trinken alle Wesen",
					"An den Brüsten der Natur,",
					"Alle Guten, alle Bösen",
					"Folgen ihrer Rosenspur.",
					"Küsse gab sie uns und Reben,",
					"Einen Freund, geprüft im Tod,",
					"Wollust ward dem Wurm gegeben,",
					"Und der Cherub steht vor Gott.",
					"Ihr stürzt nieder, Millionen?",
					"Ahnest du den Schöpfer Welt?",
					"Such ihn überm Sternenzelt,",
					"Über Sternen muß er wohnen.",
					"Freude heißt die starke Feder",
					"In der ewigen Natur.",
					"Freude, Freude treibt die Räder",
					"In der großen Weltenuhr.",
					"Blumen lockt sie aus den Keimen,",
					"Sonnen aus dem Firmament,",
					"Sphären rollt sie aus den Räumen,",
					"Die des Sehers Rohr nicht kennt.",
					"Froh, wie seine Sonnen fliegen",
					"Durch des Himmels prächt'gen Plan,",
					"Wandelt Brüder eure Bahn,",
					"Freudig wie ein Held zum Siegen.",
					"Aus der Wahrheit Feuerspiegel",
					"Lächelt sie den Forscher an.",
					"Zu der Tugend steilem Hügel",
					"Leitet sie des Dulders Bahn.",
					"Auf des Glaubens Sonnenberge",
					"Sieht man ihre Fahnen wehn,",
					"Durch den Riß gesprengter Särge",
					"Sie im Chor der Engel stehn.",
					"Duldet mutig Millionen!",
					"Duldet für die beßre Welt!",
					"Droben überm Sternenzelt",
					"Wird ein großer Gott belohnen.",
					"Göttern kann man nicht vergelten,",
					"Schön ist ihnen gleich zu sein.",
					"Gram und Armut solln sich melden,",
					"Mit den Frohen sich erfreun.",
					"Groll und Rache sein vergessen,",
					"Unserm Todfeind sei verziehn,",
					"Keine Träne soll ihn pressen,",
					"Keine Reue nage ihn.",
					"Unser Schuldbuch sei vernichtet,",
					"Ausgesöhnt die ganze Welt!",
					"Brüder – überm Sternenzelt",
					"Richtet Gott wie ihr gerichtet.",
					"Freude sprudelt in Pokalen,",
					"In der Traube goldnem Blut",
					"Trinken Sanftmut Kannibalen,",
					"Die Verzweiflung Heldenmut.",
					"Brüder fliegt von euren Sitzen,",
					"Wenn der volle Römer kreist,",
					"Laßt den Schaum zum Himmel spritzen:",
					"Dieses Glas dem guten Geist!",
					"Den der Sterne Wirbel loben,",
					"Den des Seraphs Hymne preist,",
					"Dieses Glas dem guten Geist,",
					"Überm Sternenzelt dort droben!",
					"Festen Mut in schweren Leiden,",
					"Hilfe, wo die Unschuld weint,",
					"Ewigkeit geschwornen Eiden,",
					"Wahrheit gegen Freund und Feind,",
					"Männerstolz vor Königsthronen,",
					"Brüder, gält es Gut und Blut!",
					"Dem Verdienste seine Kronen,",
					"Untergang der Lügenbrut!",
					"Schließt den heil'gen Zirkel dichter,",
					"Schwört bei diesem goldnen Wein,",
					"Dem Gelübde treu zu sein,",
					"Schwört es bei dem Sternenrichter!",
					"Rettung von Tyrannenketten,",
					"Großmut auch dem Bösewicht,",
					"Hoffnung auf den Sterbebetten,",
					"Gnade auf dem Hochgericht!",
					"Auch die Toten sollen leben!",
					"Brüder, trinkt und stimmet ein,",
					"Allen Sündern soll vergeben,",
					"Und die Hölle nicht mehr sein.",
					"Eine heitre Abschiedsstunde!",
					"Süßen Schlaf im Leichentuch!",
					"Brüder – einen sanften Spruch",
					"Aus des Totenrichters Mund.",
				}),

				Sheet: model.SheetMusic{
					Path: "symp9-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Presto",
					},
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "furelise.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "nocturne2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "nocturne20.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "sonata2-1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Grave",
					},
					{
						Form: None[string](),
						Name: "Doppio movimento",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: None[string](),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "sonata2-2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Scherzo",
					},
				},
			},
			{
				Nickname: Some("Funeral March"),
				Form: Some("Marche funèbre"),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "sonata2-3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Lento",
					},
				},
			},
			{
				Nickname: None[string](),
				Form: Some("Finale"),
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "sonata2-4.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "ballade1.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Largo",
					},
					{
						Form: None[string](),
						Name: "Moderato",
					},
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "liebestraum3.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Poco allegro, con affetto",
					},
					{
						Form: None[string](),
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
				Lyrics: None[[]string](),

				Sheet: model.SheetMusic{
					Path: "hungrhap2.mxl",
				},

				TempoMarkings: []model.TempoMarking{
					{
						Form: None[string](),
						Name: "Lento a capriccio",
					},
					{
						Form: Some("Lassan"),
						Name: "Andante mesto",
					},
					{
						Form: Some("Friska"),
						Name: "Vivace",
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
