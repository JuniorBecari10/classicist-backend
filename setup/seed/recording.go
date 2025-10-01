package seed

import "shared/model"

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
	{ Name: "Sviatoslav Richter" },
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
		Work: Works[0],
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "prelude-gould.mp3",
					Duration: 142,
				},
			},
		},
	},
	{
		Work: Works[0],
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "prelude-andras.mp3",
					Duration: 115,
				},
			},
		},
	},
	{
		Work: Works[1],
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "toccata-preston.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[1],
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "toccata-richter.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[2], // Spring
		Year: 1976,
		PhotoPath: "four-seasons-perlman.jpg",

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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "spring-1.mp3",
					Duration: 196,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "spring-2.mp3",
					Duration: 163,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "spring-3.mp3",
					Duration: 286,
				},
			},
		},
	},

	{
		Work: Works[3], // Summer
		Year: 1976,
		PhotoPath: "four-seasons-perlman.jpg",

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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "summer-1.mp3",
					Duration: 375,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "summer-2.mp3",
					Duration: 164,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "summer-3.mp3",
					Duration: 185,
				},
			},
		},
	},

	{
		Work: Works[4], // Autumn
		Year: 1976,
		PhotoPath: "four-seasons-perlman.jpg",

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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "autumn-1.mp3",
					Duration: 324,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "autumn-2.mp3",
					Duration: 186,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "autumn-3.mp3",
					Duration: 200,
				},
			},
		},
	},

	{
		Work: Works[5], // Winter
		Year: 1976,
		PhotoPath: "four-seasons-perlman.jpg",

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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "winter-1.mp3",
					Duration: 217,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "winter-2.mp3",
					Duration: 151,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "winter-3.mp3",
					Duration: 199,
				},
			},
		},
	},
	{
		Work: Works[6], // Eine kleine Nachtmusik
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "eine-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "eine-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "eine-3.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 3,
				AudioFile: model.AudioFile{
					Path: "eine-4.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[7], // Symphony No. 40
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "symp40-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "symp40-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "symp40-3.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 3,
				AudioFile: model.AudioFile{
					Path: "symp40-4.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[8], // Piano Sonata No. 16
		Year: 2023, // actually it's a live recording from 1956; it was remastered and re-released in 2023
		PhotoPath: "facile.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[8],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "facile-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "facile-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "facile-3.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[9], // Piano Sonata No. 14
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "moonlight-kempff-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "moonlight-kempff-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "moonlight-kempff-3.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[9], // Piano Sonata No. 14
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "moonlight-rubinstein-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "moonlight-rubinstein-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "moonlight-rubinstein-3.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[10], // Symphony No. 5
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "symp5-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "symp5-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "symp5-3.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 3,
				AudioFile: model.AudioFile{
					Path: "symp5-4.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[11], // Symphony No. 7
		Year: 2020,
		PhotoPath: "symp7.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[12],
				Role: "Conductor",
			},
			{
				Performer: Performers[8],
				Role: "Orchestra",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "symp7-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "symp7-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "symp7-3.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 3,
				AudioFile: model.AudioFile{
					Path: "symp7-4.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[12], // Symphony No. 9
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
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "symp9-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "symp9-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "symp9-3.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 3,
				AudioFile: model.AudioFile{
					Path: "symp9-4.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[13], // Fur Elise
		Year: 2015,
		PhotoPath: "furelise-levit.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[13],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "furelise-levit.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[13], // Fur Elise
		Year: 2025,
		PhotoPath: "furelise-ugorski.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[14],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "furelise-ugorski.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[13], // Fur Elise
		Year: 1984,
		PhotoPath: "furelise-ashkenazy.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[15],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "furelise-ashkenazy.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[14], // Nocturne No. 2
		Year: 2005,
		PhotoPath: "nocturne2-pollini.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[16],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "nocturne2-pollini.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[14], // Nocturne No. 2
		Year: 1996,
		PhotoPath: "nocturne2-maria-joao-pires.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[17],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "nocturne2-maria-joao-pires.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[15], // Nocturne No. 20
		Year: 1996,
		PhotoPath: "nocturne20.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[17],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "nocturne20.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[16], // Piano Sonata No. 2
		Year: 2023,
		PhotoPath: "sonata2.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[18],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "sonata2-1.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 1,
				AudioFile: model.AudioFile{
					Path: "sonata2-2.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 2,
				AudioFile: model.AudioFile{
					Path: "sonata2-3.mp3",
					Duration: 0xD0D0,
				},
			},
			{
				MovementIndex: 3,
				AudioFile: model.AudioFile{
					Path: "sonata2-4.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[17], // Ballade No. 1
		Year: 1988,
		PhotoPath: "ballade1-zimerman.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[19],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "ballade1-zimerman.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[17], // Ballade No. 1
		Year: 1999,
		PhotoPath: "ballade1-pollini.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[16],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "ballade1-pollini.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[18], // Liebestraum No. 3
		Year: 2011,
		PhotoPath: "liebestraum3.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[20],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "liebestraum3.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[19], // Hungarian Rhapsody No. 2
		Year: 2001,
		PhotoPath: "hungrhap2.jpg",

		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[21],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				MovementIndex: 0,
				AudioFile: model.AudioFile{
					Path: "hungrhap2.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
}
