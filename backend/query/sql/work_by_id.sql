SELECT
    -- Work
    w.id AS work_id,
    w.title_kind,
    w.title_number,
    w.title_nickname,
    w.key_note,
    w.key_mode,
    w.catalog_prefix,
    w.catalog_number,
    w.catalog_subnumber,
    w.composition_start_year,
    w.composition_end_year,

    -- Composer
    c.id AS composer_id,
    c.name AS composer_name,
    c.birth_year AS composer_birth_year,
    c.death_year AS composer_death_year,
    c.photo_path AS composer_photo_path,

    -- Movements
    m.id AS movement_id,
    m.order_num AS movement_order,
    m.kind AS movement_kind,
    m.nickname AS movement_nickname,

    -- Tempo Markings
    t.id AS tempo_id,
    t.name AS tempo_name,
    t.order_num AS tempo_order

FROM works w
JOIN composers c ON w.composer_id = c.id
LEFT JOIN movements m ON m.work_id = w.id
LEFT JOIN tempo_markings t ON t.movement_id = m.id
WHERE w.id = ?
ORDER BY m.order_num, t.order_num;
