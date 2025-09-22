SELECT
    w.*
FROM works w
JOIN composers c ON w.composer_id = c.id
LEFT JOIN recordings r ON w.id = r.work_id
WHERE w.title_kind LIKE '%' || ? || '%'
   OR c.name LIKE '%' || ? || '%'
   OR r.name LIKE '%' || ? || '%';
