SELECT r.*
FROM recordings AS r
JOIN recording_performers AS rp ON r.id = rp.recording_id
WHERE rp.performer_id = ?;
