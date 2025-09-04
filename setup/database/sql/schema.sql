CREATE TABLE IF NOT EXISTS composers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    name TEXT NOT NULL,
    birth_year INT NOT NULL,
    death_year INT,
    photo_path TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS works (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    title_kind TEXT NOT NULL,
    title_number INT,
    title_nickname TEXT,

    key_note INT NOT NULL,
    key_mode INT NOT NULL,

    composer_id INT NOT NULL,

    catalog_prefix TEXT NOT NULL,
    catalog_number TEXT,
    catalog_subnumber TEXT,

    composition_start_year INT NOT NULL,
    composition_end_year INT,

    FOREIGN KEY (composer_id) REFERENCES composers(id)
);

CREATE TABLE IF NOT EXISTS movements (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    work_id INT NOT NULL,
    order_num INT NOT NULL,
    kind TEXT,
    nickname TEXT,

    FOREIGN KEY (work_id) REFERENCES works(id)
);

CREATE TABLE IF NOT EXISTS tempo_markings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    movement_id INT NOT NULL,
    name TEXT NOT NULL,
    bpm INT,
    order_num INT NOT NULL,

    FOREIGN KEY (movement_id) REFERENCES movements(id)
);

CREATE TABLE IF NOT EXISTS recordings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    work_id INT NOT NULL,
    year INT NOT NULL,
    photo_path TEXT NOT NULL,

    FOREIGN KEY (work_id) REFERENCES works(id)
);

CREATE TABLE IF NOT EXISTS performers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS recording_performers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    recording_id INT NOT NULL,
    performer_id INT NOT NULL,
    role TEXT NOT NULL,

    FOREIGN KEY (recording_id) REFERENCES recordings(id),
    FOREIGN KEY (performer_id) REFERENCES performers(id)
);

CREATE TABLE IF NOT EXISTS recorded_movements (
    id INTEGER PRIMARY KEY AUTOINCREMENT,

    movement_id INT NOT NULL,
    recording_id INT NOT NULL,
    
	audio_path TEXT NOT NULL,
    duration INT NOT NULL,

    FOREIGN KEY (movement_id) REFERENCES movements(id),
    FOREIGN KEY (recording_id) REFERENCES recordings(id)
);
