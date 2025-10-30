package seed

import "shared/model"

var indexCount = 0
var idCount = 0
var workId = 1

func incIndex() int {
	indexCount++
	return indexCount
}

func resetIndex() int {
	indexCount = 0
	return indexCount
}

func incId() int {
	idCount++
	return idCount
}

func getWorkId() int {
	return workId
}

func incWorkId() int {
	workId++
	return workId
}

var Performers = []model.Performer{
	{ Name: "Glenn Gould" },
	{ Name: "András Schiff" },
	{ Name: "Simon Preston" },
	{ Name: "Karl Richter" },
	{ Name: "Itzhak Perlman" },
	{ Name: "Rodney Friend" },
	{ Name: "London Philharmonic Orchestra" },
	{ Name: "Herbert von Karajan" },
	{ Name: "Berliner Philharmoniker" },
	{ Name: "Wilhelm Kempff" },
	{ Name: "Arthur Rubinstein" },
	{ Name: "Kirill Petrenko" },
	{ Name: "Igor Levit" },
	{ Name: "Anatol Ugorsky" },
	{ Name: "Vladimir Ashkenazy" },
	{ Name: "Maurizio Pollini" },
	{ Name: "Maria João Pires" },
	{ Name: "Rafał Blechacz" },
	{ Name: "Krystian Zimerman" },
	{ Name: "Lang Lang" },
	{ Name: "György Cziffra" },
}

var Recordings = []model.Recording{
	{
		WorkId: getWorkId(),
		Year: 1963,
		PhotoPath: "prelude-gould.jpg",
		
		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[0],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(), // to not need to create another function just to get the id
				AudioFile: model.AudioFile{
					Path: "prelude-gould.mp3",
					Duration: 141,
				},
			},
		},
	},
	{
		WorkId: getWorkId(),
		Year: 1984,
		PhotoPath: "prelude-andras.jpg",
		
		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[1],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "prelude-andras.mp3",
					Duration: 111,
				},
			},
		},
	},
	{
		WorkId: incWorkId(),
		Year: 1988,
		PhotoPath: "toccata-preston.jpg",
		
		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[2],
				Role: "Organist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "toccata-preston.mp3",
					Duration: 499,
				},
			},
		},
	},
	{
		WorkId: getWorkId(),
		Year: 1964,
		PhotoPath: "toccata-richter.jpg",
		
		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[3],
				Role: "Organist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "toccata-richter.mp3",
					Duration: 533,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Spring
		Year: 1976,
		PhotoPath: "four-seasons.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[4],
				Role: "Violinist & Conductor",
			},
			{
				Performer: Performers[5],
				Role: "Concertmaster",
			},
			{
				Performer: Performers[6],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "spring-1.mp3",
					Duration: 189,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "spring-2.mp3",
					Duration: 156,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "spring-3.mp3",
					Duration: 280,
				},
			},
		},
	},

	{
		WorkId: incWorkId(), // Summer
		Year: 1976,
		PhotoPath: "four-seasons.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[4],
				Role: "Violinist & Conductor",
			},
			{
				Performer: Performers[5],
				Role: "Concertmaster",
			},
			{
				Performer: Performers[6],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "summer-1.mp3",
					Duration: 368,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "summer-2.mp3",
					Duration: 157,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "summer-3.mp3",
					Duration: 176,
				},
			},
		},
	},

	{
		WorkId: incWorkId(), // Autumn
		Year: 1976,
		PhotoPath: "four-seasons.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[4],
				Role: "Violinist & Conductor",
			},
			{
				Performer: Performers[5],
				Role: "Concertmaster",
			},
			{
				Performer: Performers[6],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "autumn-1.mp3",
					Duration: 316,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "autumn-2.mp3",
					Duration: 179,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "autumn-3.mp3",
					Duration: 190,
				},
			},
		},
	},

	{
		WorkId: incWorkId(), // Winter
		Year: 1976,
		PhotoPath: "four-seasons.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[4],
				Role: "Violinist & Conductor",
			},
			{
				Performer: Performers[5],
				Role: "Concertmaster",
			},
			{
				Performer: Performers[6],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "winter-1.mp3",
					Duration: 210,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "winter-2.mp3",
					Duration: 146,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "winter-3.mp3",
					Duration: 196,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Eine kleine Nachtmusik
		Year: 1966,
		PhotoPath: "eine.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[7],
				Role: "Conductor",
			},
			{
				Performer: Performers[8],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "eine-1.mp3",
					Duration: 334,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "eine-2.mp3",
					Duration: 326,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "eine-3.mp3",
					Duration: 146,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "eine-4.mp3",
					Duration: 195,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Symphony No. 40
		Year: 1977,
		PhotoPath: "symp40.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[7],
				Role: "Conductor",
			},
			{
				Performer: Performers[8],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp40-1.mp3",
					Duration: 435,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp40-2.mp3",
					Duration: 459,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp40-3.mp3",
					Duration: 274,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp40-4.mp3",
					Duration: 274,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Piano Sonata No. 16
		Year: 2023,
		PhotoPath: "facile.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[8],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "facile-1.mp3",
					Duration: 176,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "facile-2.mp3",
					Duration: 319,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "facile-3.mp3",
					Duration: 99,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Piano Sonata No. 14
		Year: 1978,
		PhotoPath: "moonlight-kempff.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[9],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "moonlight-kempff-1.mp3",
					Duration: 340,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "moonlight-kempff-2.mp3",
					Duration: 138,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "moonlight-kempff-3.mp3",
					Duration: 333,
				},
			},
		},
	},
	{
		WorkId: getWorkId(), // Piano Sonata No. 14
		Year: 1963,
		PhotoPath: "moonlight-rubinstein.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[10],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "moonlight-rubinstein-1.mp3",
					Duration: 359,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "moonlight-rubinstein-2.mp3",
					Duration: 156,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "moonlight-rubinstein-3.mp3",
					Duration: 393,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Symphony No. 5
		Year: 1984,
		PhotoPath: "symp5.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[7],
				Role: "Conductor",
			},
			{
				Performer: Performers[8],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp5-1.mp3",
					Duration: 433,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp5-2.mp3",
					Duration: 555,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp5-3.mp3",
					Duration: 285,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp5-4.mp3",
					Duration: 521,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Symphony No. 7
		Year: 2020,
		PhotoPath: "symp7.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[11],
				Role: "Conductor",
			},
			{
				Performer: Performers[8],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp7-1.mp3",
					Duration: 807,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp7-2.mp3",
					Duration: 456,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp7-3.mp3",
					Duration: 507,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp7-4.mp3",
					Duration: 478,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Symphony No. 9
		Year: 1984,
		PhotoPath: "symp9.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[7],
				Role: "Conductor",
			},
			{
				Performer: Performers[8],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp9-1.mp3",
					Duration: 917,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp9-2.mp3",
					Duration: 615,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp9-3.mp3",
					Duration: 944,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "symp9-4.mp3",
					Duration: 1451,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Fur Elise
		Year: 2015,
		PhotoPath: "furelise-levit.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[12],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "furelise-levit.mp3",
					Duration: 209,
				},
			},
		},
	},
	{
		WorkId: getWorkId(), // Fur Elise
		Year: 2025,
		PhotoPath: "furelise-ugorski.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[13],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "furelise-ugorski.mp3",
					Duration: 237,
				},
			},
		},
	},
	{
		WorkId: getWorkId(), // Fur Elise
		Year: 1984,
		PhotoPath: "furelise-ashkenazy.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[14],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "furelise-ashkenazy.mp3",
					Duration: 173,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Nocturne No. 2
		Year: 2005,
		PhotoPath: "nocturne2-pollini.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[15],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "nocturne2-pollini.mp3",
					Duration: 232,
				},
			},
		},
	},
	{
		WorkId: getWorkId(), // Nocturne No. 2
		Year: 1996,
		PhotoPath: "nocturnes-maria-joao-pires.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[16],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "nocturne2-maria-joao-pires.mp3",
					Duration: 256,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Nocturne No. 20
		Year: 1996,
		PhotoPath: "nocturnes-maria-joao-pires.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[16],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "nocturne20.mp3",
					Duration: 230,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Piano Sonata No. 2
		Year: 2023,
		PhotoPath: "sonata2.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[17],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "sonata2-1.mp3",
					Duration: 431,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "sonata2-2.mp3",
					Duration: 365,
				},
			},
			{
				MovementIndex: incIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "sonata2-3.mp3",
					Duration: 500,
				},
			},
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "sonata2-4.mp3",
					Duration: 86,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Ballade No. 1
		Year: 1988,
		PhotoPath: "ballade1-zimerman.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[18],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "ballade1-zimerman.mp3",
					Duration: 568,
				},
			},
		},
	},
	{
		WorkId: getWorkId(), // Ballade No. 1
		Year: 1999,
		PhotoPath: "ballade1-pollini.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[15],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "ballade1-pollini.mp3",
					Duration: 504,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Liebestraum No. 3
		Year: 2011,
		PhotoPath: "liebestraum3.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[19],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "liebestraum3.mp3",
					Duration: 276,
				},
			},
		},
	},
	{
		WorkId: incWorkId(), // Hungarian Rhapsody No. 2
		Year: 2001,
		PhotoPath: "hungrhap2.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[20],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: resetIndex(),
				MovementId: incId(),
				AudioFile: model.AudioFile{
					Path: "hungrhap2.mp3",
					Duration: 609,
				},
			},
		},
	},
}
