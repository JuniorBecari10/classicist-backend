SELECT t.*
FROM works w
LEFT JOIN movements m ON m.work_id = w.id
LEFT JOIN tempo_markings t ON t.movement_id = m.id
WHERE w.id = ?
ORDER BY m.order_num, t.order_num;
