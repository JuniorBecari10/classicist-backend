SELECT m.*
FROM works w
LEFT JOIN movements m ON m.work_id = w.id
WHERE w.id = ?
ORDER BY m.order_num;
