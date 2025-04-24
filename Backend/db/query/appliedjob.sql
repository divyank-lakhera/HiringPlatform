-- name: ApplyJob :one
INSERT INTO appliedjob (
    userid,
    jobid,
    recruiterid,
    current_status
) values (
    $1, $2, $3, "Applied"
) RETURNING *;

-- name: GetAllAppliedJobs :many
SELECT * FROM appliedjob
WHERE userid = $1;

-- name: UpdateJobStatus :exec
UPDATE appliedjob
SET current_status = $3
WHERE userid = $1 AND jobid = $2;

