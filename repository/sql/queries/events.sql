-- name: GetEvents :many
SELECT
    *
FROM
    events
ORDER BY
    start_time DESC
LIMIT
    $1 OFFSET $2;

-- name: GetEventByID :one
SELECT
    *
FROM
    events
WHERE
    id = $1
LIMIT
    1;

-- name: CreateEvent :one
INSERT INTO
    events (
        name,
        start_time,
        end_time,
        concurrency_target
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetTotalEventsCount :one
SELECT
    COUNT(*)
FROM
    events;