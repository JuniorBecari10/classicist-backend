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
				Kind: None[string](),
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
}
