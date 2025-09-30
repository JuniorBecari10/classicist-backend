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
}

var Recordings = []model.Recording{
	{
		Work: Works[0],
		Year: 1962,
		PhotoPath: "prelude-gould.jpg",
		
		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[0],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				Movement: Works[0].Movements[0],
				AudioFile: model.AudioFile{
					Path: "prelude-gould.mp3",
					Duration: 142,
				},
			},
		},
	},
	{
		Work: Works[0],
		Year: 0xD0D0,
		PhotoPath: "prelude-andras.jpg",
		
		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[1],
				Role: "Pianist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				Movement: Works[0].Movements[0],
				AudioFile: model.AudioFile{
					Path: "prelude-andras.mp3",
					Duration: 115,
				},
			},
		},
	},
	{
		Work: Works[1],
		Year: 0xD0D0,
		PhotoPath: "toccata-preston.jpg",
		
		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[2],
				Role: "Organist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				Movement: Works[1].Movements[0],
				AudioFile: model.AudioFile{
					Path: "toccata-preston.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[1],
		Year: 0xD0D0,
		PhotoPath: "toccata-richter.jpg",
		
		Performers: []model.RecordingPerformer{
			{
				Performer: Performers[3],
				Role: "Organist",
			},
		},

		Movements: []model.RecordedMovement{
			{
				Movement: Works[1].Movements[0],
				AudioFile: model.AudioFile{
					Path: "toccata-richter.mp3",
					Duration: 0xD0D0,
				},
			},
		},
	},
	{
		Work: Works[3], // Spring
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
				Movement: Works[3].Movements[0],
				AudioFile: model.AudioFile{
					Path: "spring-1.mp3",
					Duration: 0,
				},
			},
			{
				Movement: Works[3].Movements[1],
				AudioFile: model.AudioFile{
					Path: "spring-2.mp3",
					Duration: 0,
				},
			},
			{
				Movement: Works[3].Movements[2],
				AudioFile: model.AudioFile{
					Path: "spring-3.mp3",
					Duration: 0,
				},
			},
		},
	},

	{
		Work: Works[4], // Summer
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
				Movement: Works[4].Movements[0],
				AudioFile: model.AudioFile{
					Path: "summer-1.mp3",
					Duration: 0,
				},
			},
			{
				Movement: Works[4].Movements[1],
				AudioFile: model.AudioFile{
					Path: "summer-2.mp3",
					Duration: 0,
				},
			},
			{
				Movement: Works[4].Movements[2],
				AudioFile: model.AudioFile{
					Path: "summer-3.mp3",
					Duration: 0,
				},
			},
		},
	},

	{
		Work: Works[5], // Autumn
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
				Movement: Works[5].Movements[0],
				AudioFile: model.AudioFile{
					Path: "autumn-1.mp3",
					Duration: 0,
				},
			},
			{
				Movement: Works[5].Movements[1],
				AudioFile: model.AudioFile{
					Path: "autumn-2.mp3",
					Duration: 0,
				},
			},
			{
				Movement: Works[5].Movements[2],
				AudioFile: model.AudioFile{
					Path: "autumn-3.mp3",
					Duration: 0,
				},
			},
		},
	},

	{
		Work: Works[6], // Winter
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
				Movement: Works[6].Movements[0],
				AudioFile: model.AudioFile{
					Path: "winter-1.mp3",
					Duration: 0,
				},
			},
			{
				Movement: Works[6].Movements[1],
				AudioFile: model.AudioFile{
					Path: "winter-2.mp3",
					Duration: 0,
				},
			},
			{
				Movement: Works[6].Movements[2],
				AudioFile: model.AudioFile{
					Path: "winter-3.mp3",
					Duration: 0,
				},
			},
		},
	},
}
