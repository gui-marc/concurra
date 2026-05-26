-- name: CreateComponent :one
INSERT INTO
    components (
        name,
        team,
        area,
        type
    )
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: GetComponentByID :one
SELECT
    *
FROM
    components
WHERE
    id = $1
LIMIT
    1;

-- name: GetComponentByNameTeamArea :one
SELECT
    *
FROM
    components
WHERE
    name = $1 AND team = $2 AND area = $3
LIMIT
    1;

-- name: GetComponents :many
SELECT
    *
FROM
    components
ORDER BY
    name ASC
LIMIT
    $1 OFFSET $2;