SELECT l.*
FROM works w
LEFT JOIN movements m ON m.work_id = w.id
LEFT JOIN lyrics l ON l.movement_id = m.id
WHERE w.id = 1 AND m.id = 1
ORDER BY m.order_num, l.order_num;
