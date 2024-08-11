-- name: CreateCandidateProfile :one
INSERT INTO jobseeker (
    id,
    email,
    phone,
    firstname,
    lastname,
    photo,
    resume,
    address,
    city,
    state,
    pincode
) values (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetCandateProfile :one
SELECT * FROM jobseeker
WHERE id = $1;

-- name: UpdateCandaidatePhoto :exec
UPDATE jobseeker
SET photo = $2
WHERE id = $1;

-- name: UpdateCandaidateResume :exec
UPDATE jobseeker
SET resume = $2
WHERE id = $1;

-- name: UpdateCandidateProfile :exec
UPDATE jobseeker
SET email = $2, phone = $3, address = $4, city = $5, state = $6, pincode = $7
WHERE id = $1;