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

-- Beethoven: Symphony No. 5 in C minor ("Fate")
('Symphony', 5, 'Fate', 0, 1, 2,
 'Op.', '67', NULL, 1804, 1808),

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

-- Mahler: Symphony No. 2 in C minor
('Symphony', 2, 'Resurrection', 0, 1, 7,
 'GWV', '30', NULL, 1888, 1894),

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
INSERT INTO movements (work_id, order_num, kind) VALUES
-- Bach: Harpsichord Concerto in D minor, BWV 1052
(1, 1, NULL),
(1, 2, NULL),
(1, 3, NULL),

-- Beethoven: Piano Concerto No. 5 "Emperor"
(2, 1, NULL),
(2, 2, NULL),
(2, 3, 'Rondo'),

-- Beethoven: Piano Sonata No. 17 "Tempest"
(3, 1, NULL),
(3, 2, NULL),
(3, 3, NULL),

-- Beethoven: Symphony No. 5
(4, 1, NULL),
(4, 2, NULL),
(4, 3, 'Scherzo'),
(4, 4, 'Finale'),

-- Beethoven: Symphony No. 7
(4, 1, NULL),
(4, 2, NULL),
(4, 3, NULL),
(4, 4, NULL),

-- Beethoven: Symphony No. 9
(5, 1, NULL),
(5, 2, NULL),
(5, 3, NULL),
(5, 4, 'Finale'),

-- Chopin: Piano Sonata No. 2
(6, 1, NULL),
(6, 2, NULL),
(6, 3, 'Marche funèbre'),
(6, 4, 'Finale'),

-- Chopin: Piano Sonata No. 3
(7, 1, NULL),
(7, 2, NULL),
(7, 3, NULL),
(7, 4, NULL),

-- Chopin: Ballade No. 1
(8, 1, NULL),

-- Chopin: Ballade No. 4
(9, 1, NULL),

-- Chopin: Nocturnes
(10, 1, NULL),
(11, 1, NULL),
(12, 1, NULL),
(13, 1, NULL),
(14, 1, NULL),

-- Chopin: Piano Concerto No. 1
(15, 1, NULL),
(15, 2, NULL),
(15, 3, 'Rondo'),

-- Liszt: Liebestraum No. 3
(16, 1, 'Notturno'),

-- Mozart: Piano Concerto No. 20
(17, 1, NULL),
(17, 2, NULL),
(17, 3, 'Rondo'),

-- Mendelssohn: Symphony No. 2 "Lobgesang"
(18, 1, 'Sinfonia'),
(18, 2, NULL),
(18, 3, NULL),
(18, 4, 'Chorus'),
(18, 5, 'Chorus'),
(18, 6, 'Aria'),
(18, 7, 'Chorus'),
(18, 8, 'Aria'),
(18, 9, 'Chorus'),
(18, 10, 'Aria'),
(18, 11, 'Chorus'),
(18, 12, 'Finale'),

-- Mendelssohn: Violin Concerto
(19, 1, NULL),
(19, 2, NULL),
(19, 3, NULL),

-- Mahler: Symphony No. 2
(20, 1, NULL),
(20, 2, NULL),
(20, 3, NULL),
(20, 4, NULL),
(20, 5, 'Urlicht'),
(20, 6, 'Finale'),

-- Mahler: Symphony No. 3
(20, 1, NULL),
(20, 2, NULL),
(20, 3, NULL),
(20, 4, NULL),
(20, 5, NULL),
(20, 6, NULL),

-- Mahler: Symphony No. 5
(21, 1, NULL),
(21, 2, NULL),
(21, 3, NULL),
(21, 4, NULL),
(21, 5, NULL),

-- Mahler: Symphony No. 8
(22, 1, NULL),
(22, 2, 'Finale'),

-- Mahler: Symphony No. 10
(23, 1, NULL),
(23, 2, NULL),
(23, 3, 'Purgatorio'),
(23, 4, 'Scherzo'),
(23, 5, 'Finale'),

-- Rachmaninoff: Piano Concerto No. 2
(24, 1, NULL),
(24, 2, NULL),
(24, 3, NULL),

-- Rachmaninoff: Piano Concerto No. 3
(25, 1, NULL),
(25, 2, 'Intermezzo'),
(25, 3, 'Finale'),

-- Schubert: Symphony No. 8 "Unfinished"
(26, 1, NULL),
(26, 2, NULL);

-- Tempo markings
INSERT INTO tempo_markings (movement_id, order_num, name) VALUES
-- Bach: Harpsichord Concerto No. 1
(1, 1, 'Allegro'),
(2, 1, 'Adagio'),
(3, 1, 'Allegro'),

-- Beethoven: Piano Concerto No. 5
(4, 1, 'Allegro'),
(5, 1, 'Adagio un poco mosso'),
(6, 1, 'Allegro'),

-- Beethoven: Piano Sonata No. 17
(7, 1, 'Largo'),
(7, 2, 'Allegro'),
(8, 1, 'Adagio'),
(9, 1, 'Allegretto'),

-- Beethoven: Symphony No. 7
(10, 1, 'Allegro con brio'),
(11, 1, 'Andante con moto'),
(12, 1, 'Allegro'),
(13, 1, 'Allegro'),

-- Beethoven: Symphony No. 7
(10, 1, 'Poco sostenuto'),
(10, 2, 'Vivace'),
(11, 1, 'Allegretto'),
(12, 1, 'Presto'),
(12, 2, 'Assai meno presto'),
(13, 1, 'Allegro con brio'),

-- Beethoven: Symphony No. 9
(14, 1, 'Allegro ma non troppo, un poco maestoso'),
(15, 1, 'Molto vivace'),
(16, 1, 'Adagio molto e cantabile'),
(17, 1, 'Presto'),
(17, 2, 'Allegro assai'),

-- Chopin: Piano Sonata No. 2
(18, 1, 'Grave'),
(18, 2, 'Doppio movimento'),
(19, 1, 'Scherzo'),
(20, 1, 'Lento'),
(21, 1, 'Presto'),

-- Chopin: Piano Sonata No. 3
(22, 1, 'Allegro maestoso'),
(23, 1, 'Molto vivace'),
(24, 1, 'Largo'),
(25, 1, 'Presto non tanto'),

-- Chopin: Ballade No. 1
(26, 1, 'Largo'),
(26, 2, 'Moderato'),
(26, 3, 'Presto con fuoco'),

-- Chopin: Ballade No. 4
(27, 1, 'Andante con moto'),

-- Chopin: Nocturnes
(28, 1, 'Larghetto'),
(29, 1, 'Andante'),
(30, 1, 'Andante cantabile'),
(31, 1, 'Lento sostenuto'),
(32, 1, 'Lento con gran espressione'),

-- Chopin: Piano Concerto No. 1
(33, 1, 'Allegro maestoso'),
(34, 1, 'Larghetto'),
(35, 1, 'Vivace'),

-- Liszt: Liebestraum No. 3
(36, 1, 'Poco allegro, con affetto'),
(36, 2, 'Tempo I'),

-- Mozart: Piano Concerto No. 20
(37, 1, 'Allegro'),
(38, 1, 'Romanze'),
(39, 1, 'Allegro assai'),

-- Mendelssohn: Symphony No. 2 "Lobgesang"
(40, 1, 'Maestoso con moto'),
(41, 1, 'Allegretto un poco agitato'),
(42, 1, 'Adagio religioso'),
(43, 1, 'O Herr, lehre mich'),
(44, 1, 'Nun danket alle Gott'),
(45, 1, 'Ihr Sorgen, flieht'),
(46, 1, 'Wer bis an das Ende beharrt'),
(47, 1, 'Ihr Völker des Himmels'),
(48, 1, 'Halleluja'),
(49, 1, 'Preiset den Herrn'),
(50, 1, 'Lobet den Herrn'),
(51, 1, 'Alles was Odem hat, lobe den Herrn'),

-- Mendelssohn: Violin Concerto
(52, 1, 'Allegro molto appassionato'),
(53, 1, 'Andante'),
(54, 1, 'Allegretto non troppo'),
(54, 2, 'Allegro molto vivace'),

-- Mahler: Symphony No. 3
(55, 1, 'Kräftig. Entschieden'),
(56, 1, 'Tempo di Menuetto'),
(57, 1, 'Comodo. Scherzando'),
(58, 1, 'Sehr langsam'),
(58, 2, 'Misterioso'),
(59, 1, 'Lustig im Tempo und keck im Ausdruck'),
(60, 1, 'Langsam. Ruhevoll. Empfunden'),

-- Mahler: Symphony No. 5
(61, 1, 'Trauermarsch'),
(62, 1, 'Stürmisch bewegt'),
(63, 1, 'Scherzo'),
(64, 1, 'Adagietto'),
(65, 1, 'Rondo-Finale'),

-- Mahler: Symphony No. 8
(66, 1, 'Veni, creator spiritus'),
(67, 1, 'Chorus mysticus'),

-- Mahler: Symphony No. 10
(68, 1, 'Adagio'),
(69, 1, 'Scherzo'),
(70, 1, 'Allegretto moderato'),
(71, 1, 'Allegro pesante'),
(72, 1, 'Langsam, schwer'),

-- Rachmaninoff: Piano Concerto No. 2
(73, 1, 'Moderato'),
(74, 1, 'Adagio sostenuto'),
(75, 1, 'Allegro scherzando'),

-- Rachmaninoff: Piano Concerto No. 3
(76, 1, 'Allegro ma non tanto'),
(77, 1, 'Adagio'),
(78, 1, 'Alla breve'),

-- Schubert: Symphony No. 8 "Unfinished"
(79, 1, 'Allegro moderato'),
(80, 1, 'Andante con moto');
