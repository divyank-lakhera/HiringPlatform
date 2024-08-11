// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: appliedjob.sql

package db

import (
	"context"
)

const applyJob = `-- name: ApplyJob :one
INSERT INTO appliedjob (
    userid,
    jobid,
    recruiterid,
    current_status
) values (
    $1, $2, $3, "Applied"
) RETURNING userid, jobid, recruiterid, current_status
`

type ApplyJobParams struct {
	Userid      int64 `json:"userid"`
	Jobid       int64 `json:"jobid"`
	Recruiterid int64 `json:"recruiterid"`
}

func (q *Queries) ApplyJob(ctx context.Context, arg ApplyJobParams) (Appliedjob, error) {
	row := q.db.QueryRowContext(ctx, applyJob, arg.Userid, arg.Jobid, arg.Recruiterid)
	var i Appliedjob
	err := row.Scan(
		&i.Userid,
		&i.Jobid,
		&i.Recruiterid,
		&i.CurrentStatus,
	)
	return i, err
}

const getAllAppliedJobs = `-- name: GetAllAppliedJobs :many
SELECT userid, jobid, recruiterid, current_status FROM appliedjob
WHERE userid = $1
`

func (q *Queries) GetAllAppliedJobs(ctx context.Context, userid int64) ([]Appliedjob, error) {
	rows, err := q.db.QueryContext(ctx, getAllAppliedJobs, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Appliedjob
	for rows.Next() {
		var i Appliedjob
		if err := rows.Scan(
			&i.Userid,
			&i.Jobid,
			&i.Recruiterid,
			&i.CurrentStatus,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateJobStatus = `-- name: UpdateJobStatus :exec
UPDATE appliedjob
SET current_status = $3
WHERE userid = $1 AND jobid = $2
`

type UpdateJobStatusParams struct {
	Userid        int64  `json:"userid"`
	Jobid         int64  `json:"jobid"`
	CurrentStatus string `json:"current_status"`
}

func (q *Queries) UpdateJobStatus(ctx context.Context, arg UpdateJobStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateJobStatus, arg.Userid, arg.Jobid, arg.CurrentStatus)
	return err
}
