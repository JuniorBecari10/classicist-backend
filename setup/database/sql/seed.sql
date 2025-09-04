-- Insert composers
INSERT INTO composers (name, birth_year, death_year, photo_path) VALUES
('Johann Sebastian Bach', 1685, 1750, 'bach.jpg'),
('Ludwig van Beethoven', 1770, 1827, 'beethoven.jpg'),
('Frédéric Chopin', 1810, 1849, 'chopin.jpg'),
('Franz Liszt', 1811, 1886, 'liszt.jpg'),
('Gustav Mahler', 1860, 1911, 'mahler.jpg'),
('Felix Mendelssohn', 1809, 1847, 'mendelssohn.jpg'),
('Wolfgang Amadeus Mozart', 1756, 1791, 'mozart.jpg'),
('Sergei Rachmaninoff', 1873, 1943, 'rachmaninoff.jpg'),
('Franz Schubert', 1797, 1828, 'schubert.jpg');

-- Insert works
INSERT INTO works (
    title_kind, title_number, title_nickname,
    key_note, key_mode, composer_id,
    catalog_prefix, catalog_number, catalog_subnumber,
    composition_start_year, composition_end_year
) VALUES
-- Bach: Harpsichord Concerto in D minor
('Harpsichord Concerto', 1, NULL, 3, 1, 1,
 'BWV', '1052', NULL, 1730, NULL),

-- Beethoven: Piano Concerto No. 5 in E-flat major
('Piano Concerto', 5, 'Emperor', 5, 0, 2,
 'Op.', '73', NULL, 1809, NULL),

-- Beethoven: Piano Sonata No. 17 in D minor ("Tempest")
('Piano Sonata', 17, 'Tempest', 3, 1, 2,
 'Op.', '31', '2', 1801, NULL),

-- Beethoven: Symphony No. 7 in A major
('Symphony', 7, NULL, 13, 0, 2,
 'Op.', '92', NULL, 1811, NULL),

-- Beethoven: Symphony No. 9 in D minor
('Symphony', 9, 'Choral', 3, 1, 2,
 'Op.', '125', NULL, 1822, 1824),

-- Chopin: Piano Sonata No. 2 in B-flat minor
('Piano Sonata', 2, NULL, 15, 1, 3,
 'Op.', '35', NULL, 1839, 1840),

-- Chopin: Piano Sonata No. 3 in B minor
('Piano Sonata', 3, NULL, 16, 1, 3,
 'Op.', '58', NULL, 1844, NULL),

-- Chopin: Ballade No. 1 in G minor
('Ballade', 1, NULL, 10, 1, 3,
 'Op.', '23', NULL, 1831, 1835),

-- Chopin: Ballade No. 4 in F minor
('Ballade', 4, NULL, 7, 1, 3,
 'Op.', '52', NULL, 1842, NULL),

-- Chopin: Nocturne Op. 9 No. 1 in B-flat minor
('Nocturne', 1, NULL, 15, 1, 3,
 'Op.', '9', '1', 1830, 1832),

-- Chopin: Nocturne Op. 9 No. 2 in E-flat major
('Nocturne', 2, NULL, 5, 0, 3,
 'Op.', '9', '2', 1830, 1832),

-- Chopin: Nocturne Op. 15 No. 1 in F major
('Nocturne', 1, NULL, 7, 0, 3,
 'Op.', '15', '1', 1830, 1833),

-- Chopin: Nocturne Op. 27 No. 2 in D-flat major
('Nocturne', 2, NULL, 2, 0, 3,
 'Op.', '27', '2', 1835, 1836),

-- Chopin: Nocturne in C-sharp minor (posth.)
('Nocturne', 20, NULL, 1, 1, 3,
 'Op. posth.', NULL, NULL, 1830, NULL),

-- Chopin: Piano Concerto No. 1 in E minor
('Piano Concerto', 1, NULL, 6, 1, 3,
 'Op.', '11', NULL, 1830, NULL),

-- Liszt: Liebestraum No. 3 in A-flat major
('Liebestraum', 3, NULL, 12, 0, 4,
 'S.', '541', '3', 1850, NULL),

-- Mozart: Piano Concerto No. 20 in D minor
('Piano Concerto', 20, NULL, 3, 1, 6,
 'K.', '466', NULL, 1785, NULL),

-- Mendelssohn: Symphony No. 2 ("Lobgesang") in B-flat major
('Symphony', 2, 'Lobgesang', 15, 0, 5,
 'Op.', '52', NULL, 1840, NULL),

-- Mendelssohn: Violin Concerto in E minor
('Violin Concerto', NULL, NULL, 6, 1, 5,
 'Op.', '64', NULL, 1844, NULL),

-- Mahler: Symphony No. 3 in D minor
('Symphony', 3, NULL, 3, 1, 7,
 'GWV', '31', NULL, 1893, 1896),

-- Mahler: Symphony No. 5 in C-sharp minor
('Symphony', 5, NULL, 1, 1, 7,
 'GWV', '44', NULL, 1901, 1902),

-- Mahler: Symphony No. 8 in E-flat major
('Symphony', 8, 'Symphony of a Thousand', 5, 0, 7,
 'GWV', '48', NULL, 1906, 1907),

-- Mahler: Symphony No. 10 in F-sharp major
('Symphony', 10, NULL, 8, 0, 7,
 'GWV', 'F58', NULL, 1910, NULL),

-- Rachmaninoff: Piano Concerto No. 2 in C minor
('Piano Concerto', 2, NULL, 0, 1, 8,
 'Op.', '18', NULL, 1900, 1901),

-- Rachmaninoff: Piano Concerto No. 3 in D minor
('Piano Concerto', 3, NULL, 3, 1, 8,
 'Op.', '30', NULL, 1909, NULL),

-- Schubert: Symphony No. 8 in B minor ("Unfinished")
('Symphony', 8, 'Unfinished', 16, 1, 9,
 'D.', '759', NULL, 1822, NULL);

-- Movements
INSERT INTO movements (work_id, order_num, kind, nickname) VALUES
-- Bach: Harpsichord Concerto in D minor, BWV 1052
(1, 1, NULL, NULL),
(1, 2, NULL, NULL),
(1, 3, NULL, NULL),

-- Beethoven: Piano Concerto No. 5 "Emperor"
(2, 1, NULL, NULL),
(2, 2, NULL, NULL),
(2, 3, 'Rondo', NULL),

-- Beethoven: Piano Sonata No. 17 "Tempest"
(3, 1, NULL, NULL),
(3, 2, NULL, NULL),
(3, 3, NULL, NULL),

-- Beethoven: Symphony No. 7
(4, 1, NULL, NULL),
(4, 2, NULL, NULL),
(4, 3, NULL, NULL),
(4, 4, NULL, NULL),

-- Beethoven: Symphony No. 9
(5, 1, NULL, NULL),
(5, 2, NULL, NULL),
(5, 3, NULL, NULL),
(5, 4, 'Finale', 'Ode to Joy'),

-- Chopin: Piano Sonata No. 2
(6, 1, NULL, NULL),
(6, 2, NULL, NULL),
(6, 3, 'Marche funèbre', NULL),
(6, 4, 'Finale', NULL),

-- Chopin: Piano Sonata No. 3
(7, 1, NULL, NULL),
(7, 2, NULL, NULL),
(7, 3, NULL, NULL),
(7, 4, NULL, NULL),

-- Chopin: Ballade No. 1
(8, 1, NULL, NULL),

-- Chopin: Ballade No. 4
(9, 1, NULL, NULL),

-- Chopin: Nocturnes
(10, 1, NULL, NULL),
(11, 1, NULL, NULL),
(12, 1, NULL, NULL),
(13, 1, NULL, NULL),
(14, 1, NULL, NULL),

-- Chopin: Piano Concerto No. 1
(15, 1, NULL, NULL),
(15, 2, NULL, NULL),
(15, 3, 'Rondo', NULL),

-- Liszt: Liebestraum No. 3
(16, 1, 'Notturno', NULL),

-- Mozart: Piano Concerto No. 20
(17, 1, NULL, NULL),
(17, 2, NULL, NULL),
(17, 3, 'Rondo', NULL),

-- Mendelssohn: Symphony No. 2 "Lobgesang"
(18, 1, 'Sinfonia', NULL),
(18, 2, NULL, NULL),
(18, 3, NULL, NULL),
(18, 4, 'Chorale', NULL),
(18, 5, 'Chorus', NULL),
(18, 6, 'Aria', NULL),
(18, 7, 'Chorus', NULL),
(18, 8, 'Aria', NULL),
(18, 9, 'Chorus', NULL),
(18, 10, 'Aria', NULL),
(18, 11, 'Chorus', NULL),
(18, 12, 'Finale', NULL),

-- Mendelssohn: Violin Concerto
(19, 1, NULL, NULL),
(19, 2, NULL, NULL),
(19, 3, NULL, NULL),

-- Mahler: Symphony No. 3
(20, 1, NULL, NULL),
(20, 2, NULL, NULL),
(20, 3, NULL, NULL),
(20, 4, NULL, NULL),
(20, 5, NULL, NULL),
(20, 6, NULL, NULL),

-- Mahler: Symphony No. 5
(21, 1, NULL, NULL),
(21, 2, NULL, NULL),
(21, 3, NULL, NULL),
(21, 4, NULL, NULL),
(21, 5, NULL, NULL),

-- Mahler: Symphony No. 8
(22, 1, NULL, NULL),
(22, 2, 'Finale', NULL),

-- Mahler: Symphony No. 10
(23, 1, NULL, NULL),
(23, 2, NULL, NULL),
(23, 3, 'Purgatorio', NULL),
(23, 4, 'Scherzo', NULL),
(23, 5, 'Finale', NULL),

-- Rachmaninoff: Piano Concerto No. 2
(24, 1, NULL, NULL),
(24, 2, NULL, NULL),
(24, 3, NULL, NULL),

-- Rachmaninoff: Piano Concerto No. 3
(25, 1, NULL, NULL),
(25, 2, 'Intermezzo', NULL),
(25, 3, 'Finale', NULL),

-- Schubert: Symphony No. 8 "Unfinished"
(26, 1, NULL, NULL),
(26, 2, NULL, NULL);

-- Tempo markings
INSERT INTO tempo_markings (movement_id, order_num, name, bpm) VALUES
-- Bach
(1, 1, 'Allegro', NULL),
(2, 1, 'Adagio', NULL),
(3, 1, 'Allegro', NULL),

-- Beethoven: Piano Concerto No. 5
(4, 1, 'Allegro', NULL),
(5, 1, 'Adagio un poco mosso', NULL),
(6, 1, 'Rondo: Allegro', NULL),

-- Beethoven: Piano Sonata No. 17
(7, 1, 'Largo', NULL),
(7, 2, 'Allegro', NULL),
(8, 1, 'Adagio', NULL),
(9, 1, 'Allegretto', NULL),

-- Beethoven: Symphony No. 7
(10, 1, 'Poco sostenuto', NULL),
(10, 2, 'Vivace', NULL),
(11, 1, 'Allegretto', NULL),
(12, 1, 'Presto', NULL),
(12, 2, 'Assai meno presto', NULL),
(13, 1, 'Allegro con brio', NULL),

-- Beethoven: Symphony No. 9
(14, 1, 'Allegro ma non troppo, un poco maestoso', NULL),
(15, 1, 'Molto vivace', NULL),
(16, 1, 'Adagio molto e cantabile', NULL),
(17, 1, 'Presto', NULL),
(17, 2, 'Allegro assai', NULL),

-- Chopin: Piano Sonata No. 2
(18, 1, 'Grave', NULL),
(18, 2, 'Doppio movimento', NULL),
(19, 1, 'Scherzo', NULL),
(20, 1, 'Lento', NULL),
(21, 1, 'Presto', NULL),

-- Chopin: Piano Sonata No. 3
(22, 1, 'Allegro maestoso', NULL),
(23, 1, 'Molto vivace', NULL),
(24, 1, 'Largo', NULL),
(25, 1, 'Presto non tanto', NULL),

-- Chopin: Ballade No. 1
(26, 1, 'Largo', NULL),
(26, 2, 'Moderato', NULL),
(26, 3, 'Presto con fuoco', NULL),

-- Chopin: Ballade No. 4
(27, 1, 'Andante con moto', NULL),

-- Chopin: Nocturnes
(28, 1, 'Larghetto', NULL),
(29, 1, 'Andante', NULL),
(30, 1, 'Andante cantabile', NULL),
(31, 1, 'Lento sostenuto', NULL),
(32, 1, 'Lento con gran espressione', NULL);

-- Chopin: Piano Concerto No. 1
(33, 1, 'Allegro maestoso', NULL),
(34, 1, 'Larghetto', NULL),
(35, 1, 'Vivace', NULL),

-- Liszt: Liebestraum No. 3
(36, 1, 'Poco allegro, con affetto', NULL),
(36, 2, 'Tempo I', NULL),

-- Mozart: Piano Concerto No. 20
(37, 1, 'Allegro', NULL),
(38, 1, 'Romanze', NULL),
(39, 1, 'Allegro assai', NULL),

-- Mendelssohn: Symphony No. 2 "Lobgesang"
(40, 1, 'Maestoso con moto', NULL),
(41, 1, 'Allegretto un poco agitato', NULL),
(42, 1, 'Adagio religioso', NULL),
(43, 1, 'O Herr, lehre mich', NULL),
(44, 1, 'Nun danket alle Gott', NULL),
(45, 1, 'Ihr Sorgen, flieht', NULL),
(46, 1, 'Wer bis an das Ende beharrt', NULL),
(47, 1, 'Ihr Völker des Himmels', NULL),
(48, 1, 'Halleluja', NULL),
(49, 1, 'Preiset den Herrn', NULL),
(50, 1, 'Lobet den Herrn', NULL),
(51, 1, 'Alles was Odem hat, lobe den Herrn', NULL),

-- Mendelssohn: Violin Concerto
(52, 1, 'Allegro molto appassionato', NULL),
(53, 1, 'Andante', NULL),
(54, 1, 'Allegretto non troppo', NULL),
(54, 2, 'Allegro molto vivace', NULL),

-- Mahler: Symphony No. 3
(55, 1, 'Kräftig. Entschieden', NULL),
(56, 1, 'Tempo di Menuetto', NULL),
(57, 1, 'Comodo. Scherzando', NULL),
(58, 1, 'Sehr langsam', NULL),
(58, 2, 'Misterioso', NULL),
(59, 1, 'Lustig im Tempo und keck im Ausdruck', NULL),
(60, 1, 'Langsam. Ruhevoll. Empfunden', NULL),

-- Mahler: Symphony No. 5
(61, 1, 'Trauermarsch', NULL),
(62, 1, 'Stürmisch bewegt', NULL),
(63, 1, 'Scherzo', NULL),
(64, 1, 'Adagietto', NULL),
(65, 1, 'Rondo-Finale', NULL),

-- Mahler: Symphony No. 8
(66, 1, 'Veni, creator spiritus', NULL),
(67, 1, 'Chorus mysticus', NULL),

-- Mahler: Symphony No. 10
(68, 1, 'Adagio', NULL),
(69, 1, 'Scherzo', NULL),
(70, 1, 'Allegretto moderato', NULL),
(71, 1, 'Allegro pesante', NULL),
(72, 1, 'Langsam, schwer', NULL),

-- Rachmaninoff: Piano Concerto No. 2
(73, 1, 'Moderato', NULL),
(74, 1, 'Adagio sostenuto', NULL),
(75, 1, 'Allegro scherzando', NULL),

-- Rachmaninoff: Piano Concerto No. 3
(76, 1, 'Allegro ma non tanto', NULL),
(77, 1, 'Adagio', NULL),
(78, 1, 'Alla breve', NULL),

-- Schubert: Symphony No. 8 "Unfinished"
(79, 1, 'Allegro moderato', NULL),
(80, 1, 'Andante con moto', NULL);
