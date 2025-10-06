SELECT l.*
FROM works w
JOIN movements m ON m.work_id = w.id
JOIN lyrics l ON l.movement_id = m.id
WHERE w.id = ? AND m.id = ?
ORDER BY m.order_num, l.order_num;
