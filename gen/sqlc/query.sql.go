// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package autoredial

import (
	"context"
	"database/sql"
	"time"
)

const createCallHistory = `-- name: CreateCallHistory :one
INSERT INTO callhistory (
  id, scammer_id, call_at, call_time, result, talk_time
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, scammer_id, call_at, call_time, result, talk_time
`

type CreateCallHistoryParams struct {
	ID        string
	ScammerID string
	CallAt    time.Time
	CallTime  int64
	Result    bool
	TalkTime  sql.NullInt64
}

func (q *Queries) CreateCallHistory(ctx context.Context, arg CreateCallHistoryParams) (Callhistory, error) {
	row := q.db.QueryRowContext(ctx, createCallHistory,
		arg.ID,
		arg.ScammerID,
		arg.CallAt,
		arg.CallTime,
		arg.Result,
		arg.TalkTime,
	)
	var i Callhistory
	err := row.Scan(
		&i.ID,
		&i.ScammerID,
		&i.CallAt,
		&i.CallTime,
		&i.Result,
		&i.TalkTime,
	)
	return i, err
}

const createMatching = `-- name: CreateMatching :one
INSERT INTO matching (
  id, created_at, serial_number, matched, checked, matching_at, talk_time, transcript
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, created_at, serial_number, matched, checked, matching_at, talk_time, transcript
`

type CreateMatchingParams struct {
	ID           string
	CreatedAt    time.Time
	SerialNumber int64
	Matched      bool
	Checked      bool
	MatchingAt   sql.NullTime
	TalkTime     sql.NullInt64
	Transcript   sql.NullString
}

func (q *Queries) CreateMatching(ctx context.Context, arg CreateMatchingParams) (Matching, error) {
	row := q.db.QueryRowContext(ctx, createMatching,
		arg.ID,
		arg.CreatedAt,
		arg.SerialNumber,
		arg.Matched,
		arg.Checked,
		arg.MatchingAt,
		arg.TalkTime,
		arg.Transcript,
	)
	var i Matching
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.SerialNumber,
		&i.Matched,
		&i.Checked,
		&i.MatchingAt,
		&i.TalkTime,
		&i.Transcript,
	)
	return i, err
}

const createMetric = `-- name: CreateMetric :one
INSERT INTO metric (
  id, created_at, calls, scammers, inactives, call_time, talk_time, amount
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, created_at, calls, scammers, inactives, call_time, talk_time, amount
`

type CreateMetricParams struct {
	ID        string
	CreatedAt time.Time
	Calls     int64
	Scammers  int64
	Inactives int64
	CallTime  int64
	TalkTime  int64
	Amount    int64
}

func (q *Queries) CreateMetric(ctx context.Context, arg CreateMetricParams) (Metric, error) {
	row := q.db.QueryRowContext(ctx, createMetric,
		arg.ID,
		arg.CreatedAt,
		arg.Calls,
		arg.Scammers,
		arg.Inactives,
		arg.CallTime,
		arg.TalkTime,
		arg.Amount,
	)
	var i Metric
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Calls,
		&i.Scammers,
		&i.Inactives,
		&i.CallTime,
		&i.TalkTime,
		&i.Amount,
	)
	return i, err
}

const createScammer = `-- name: CreateScammer :one
INSERT INTO scammer (
  id, name, tel, is_active
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, tel, is_active
`

type CreateScammerParams struct {
	ID       string
	Name     string
	Tel      string
	IsActive bool
}

func (q *Queries) CreateScammer(ctx context.Context, arg CreateScammerParams) (Scammer, error) {
	row := q.db.QueryRowContext(ctx, createScammer,
		arg.ID,
		arg.Name,
		arg.Tel,
		arg.IsActive,
	)
	var i Scammer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Tel,
		&i.IsActive,
	)
	return i, err
}

const createScammerCall = `-- name: CreateScammerCall :one
INSERT INTO scammer_call (
  scammer_id, call_id
) VALUES (
  $1, $2
)
RETURNING scammer_id, call_id
`

type CreateScammerCallParams struct {
	ScammerID string
	CallID    string
}

func (q *Queries) CreateScammerCall(ctx context.Context, arg CreateScammerCallParams) (ScammerCall, error) {
	row := q.db.QueryRowContext(ctx, createScammerCall, arg.ScammerID, arg.CallID)
	var i ScammerCall
	err := row.Scan(&i.ScammerID, &i.CallID)
	return i, err
}

const createScammerMatching = `-- name: CreateScammerMatching :one
INSERT INTO scammer_matching (
  scammer_id, matching_id
) VALUES (
  $1, $2
)
RETURNING scammer_id, matching_id
`

type CreateScammerMatchingParams struct {
	ScammerID  string
	MatchingID string
}

func (q *Queries) CreateScammerMatching(ctx context.Context, arg CreateScammerMatchingParams) (ScammerMatching, error) {
	row := q.db.QueryRowContext(ctx, createScammerMatching, arg.ScammerID, arg.MatchingID)
	var i ScammerMatching
	err := row.Scan(&i.ScammerID, &i.MatchingID)
	return i, err
}

const createScammerTag = `-- name: CreateScammerTag :one
INSERT INTO scammer_tag (
  scammer_id, tag
) VALUES (
  $1, $2
)
RETURNING scammer_id, tag
`

type CreateScammerTagParams struct {
	ScammerID string
	Tag       string
}

func (q *Queries) CreateScammerTag(ctx context.Context, arg CreateScammerTagParams) (ScammerTag, error) {
	row := q.db.QueryRowContext(ctx, createScammerTag, arg.ScammerID, arg.Tag)
	var i ScammerTag
	err := row.Scan(&i.ScammerID, &i.Tag)
	return i, err
}

const deleteCallHistory = `-- name: DeleteCallHistory :exec
DELETE FROM callhistory
WHERE id = $1
`

func (q *Queries) DeleteCallHistory(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCallHistory, id)
	return err
}

const deleteMatching = `-- name: DeleteMatching :exec
DELETE FROM matching
WHERE id = $1
`

func (q *Queries) DeleteMatching(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteMatching, id)
	return err
}

const deleteMetric = `-- name: DeleteMetric :exec
DELETE FROM metric
WHERE created_at = $1
`

func (q *Queries) DeleteMetric(ctx context.Context, createdAt time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteMetric, createdAt)
	return err
}

const deleteScammer = `-- name: DeleteScammer :exec
DELETE FROM scammer
WHERE id = $1
`

func (q *Queries) DeleteScammer(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteScammer, id)
	return err
}

const deleteScammerCall = `-- name: DeleteScammerCall :exec
DELETE FROM scammer_call
WHERE scammer_id = $1 AND call_id = $2
`

type DeleteScammerCallParams struct {
	ScammerID string
	CallID    string
}

func (q *Queries) DeleteScammerCall(ctx context.Context, arg DeleteScammerCallParams) error {
	_, err := q.db.ExecContext(ctx, deleteScammerCall, arg.ScammerID, arg.CallID)
	return err
}

const deleteScammerMatching = `-- name: DeleteScammerMatching :exec
DELETE FROM scammer_matching
WHERE scammer_id = $1 AND matching_id = $2
`

type DeleteScammerMatchingParams struct {
	ScammerID  string
	MatchingID string
}

func (q *Queries) DeleteScammerMatching(ctx context.Context, arg DeleteScammerMatchingParams) error {
	_, err := q.db.ExecContext(ctx, deleteScammerMatching, arg.ScammerID, arg.MatchingID)
	return err
}

const deleteScammerTag = `-- name: DeleteScammerTag :exec
DELETE FROM scammer_tag
WHERE scammer_id = $1 AND tag = $2
`

type DeleteScammerTagParams struct {
	ScammerID string
	Tag       string
}

func (q *Queries) DeleteScammerTag(ctx context.Context, arg DeleteScammerTagParams) error {
	_, err := q.db.ExecContext(ctx, deleteScammerTag, arg.ScammerID, arg.Tag)
	return err
}

const getCallHistory = `-- name: GetCallHistory :one
SELECT id, scammer_id, call_at, call_time, result, talk_time FROM callhistory
WHERE id = $1
`

func (q *Queries) GetCallHistory(ctx context.Context, id string) (Callhistory, error) {
	row := q.db.QueryRowContext(ctx, getCallHistory, id)
	var i Callhistory
	err := row.Scan(
		&i.ID,
		&i.ScammerID,
		&i.CallAt,
		&i.CallTime,
		&i.Result,
		&i.TalkTime,
	)
	return i, err
}

const getScammer = `-- name: GetScammer :one
SELECT id, name, tel, is_active FROM scammer
WHERE id = $1
`

func (q *Queries) GetScammer(ctx context.Context, id string) (Scammer, error) {
	row := q.db.QueryRowContext(ctx, getScammer, id)
	var i Scammer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Tel,
		&i.IsActive,
	)
	return i, err
}

const listCallHistoryByScammerId = `-- name: ListCallHistoryByScammerId :many
SELECT id, scammer_id, call_at, call_time, result, talk_time FROM callhistory
WHERE scammer_id = $1
`

func (q *Queries) ListCallHistoryByScammerId(ctx context.Context, scammerID string) ([]Callhistory, error) {
	rows, err := q.db.QueryContext(ctx, listCallHistoryByScammerId, scammerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Callhistory
	for rows.Next() {
		var i Callhistory
		if err := rows.Scan(
			&i.ID,
			&i.ScammerID,
			&i.CallAt,
			&i.CallTime,
			&i.Result,
			&i.TalkTime,
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

const listMatching = `-- name: ListMatching :many
SELECT id, created_at, serial_number, matched, checked, matching_at, talk_time, transcript FROM matching
WHERE id = $1
`

func (q *Queries) ListMatching(ctx context.Context, id string) ([]Matching, error) {
	rows, err := q.db.QueryContext(ctx, listMatching, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Matching
	for rows.Next() {
		var i Matching
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.SerialNumber,
			&i.Matched,
			&i.Checked,
			&i.MatchingAt,
			&i.TalkTime,
			&i.Transcript,
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

const listMetric = `-- name: ListMetric :many
SELECT id, created_at, calls, scammers, inactives, call_time, talk_time, amount FROM metric
WHERE created_at BETWEEN $1 AND $2
ORDER BY created_at DESC
`

type ListMetricParams struct {
	CreatedAt   time.Time
	CreatedAt_2 time.Time
}

func (q *Queries) ListMetric(ctx context.Context, arg ListMetricParams) ([]Metric, error) {
	rows, err := q.db.QueryContext(ctx, listMetric, arg.CreatedAt, arg.CreatedAt_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Metric
	for rows.Next() {
		var i Metric
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Calls,
			&i.Scammers,
			&i.Inactives,
			&i.CallTime,
			&i.TalkTime,
			&i.Amount,
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

const listScammerCallByCallId = `-- name: ListScammerCallByCallId :many
SELECT scammer_id, call_id FROM scammer_call
WHERE call_id = $1
`

func (q *Queries) ListScammerCallByCallId(ctx context.Context, callID string) ([]ScammerCall, error) {
	rows, err := q.db.QueryContext(ctx, listScammerCallByCallId, callID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScammerCall
	for rows.Next() {
		var i ScammerCall
		if err := rows.Scan(&i.ScammerID, &i.CallID); err != nil {
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

const listScammerCallByScammerId = `-- name: ListScammerCallByScammerId :many
SELECT scammer_id, call_id FROM scammer_call
WHERE scammer_id = $1
`

func (q *Queries) ListScammerCallByScammerId(ctx context.Context, scammerID string) ([]ScammerCall, error) {
	rows, err := q.db.QueryContext(ctx, listScammerCallByScammerId, scammerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScammerCall
	for rows.Next() {
		var i ScammerCall
		if err := rows.Scan(&i.ScammerID, &i.CallID); err != nil {
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

const listScammerMatchingByMatchingId = `-- name: ListScammerMatchingByMatchingId :many
SELECT scammer_id, matching_id FROM scammer_matching
WHERE matching_id = $1
`

func (q *Queries) ListScammerMatchingByMatchingId(ctx context.Context, matchingID string) ([]ScammerMatching, error) {
	rows, err := q.db.QueryContext(ctx, listScammerMatchingByMatchingId, matchingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScammerMatching
	for rows.Next() {
		var i ScammerMatching
		if err := rows.Scan(&i.ScammerID, &i.MatchingID); err != nil {
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

const listScammerMatchingByScammerId = `-- name: ListScammerMatchingByScammerId :many
SELECT scammer_id, matching_id FROM scammer_matching
WHERE scammer_id = $1
`

func (q *Queries) ListScammerMatchingByScammerId(ctx context.Context, scammerID string) ([]ScammerMatching, error) {
	rows, err := q.db.QueryContext(ctx, listScammerMatchingByScammerId, scammerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScammerMatching
	for rows.Next() {
		var i ScammerMatching
		if err := rows.Scan(&i.ScammerID, &i.MatchingID); err != nil {
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

const listScammerTag = `-- name: ListScammerTag :many
SELECT scammer_id, tag FROM scammer_tag
WHERE scammer_id = $1
`

func (q *Queries) ListScammerTag(ctx context.Context, scammerID string) ([]ScammerTag, error) {
	rows, err := q.db.QueryContext(ctx, listScammerTag, scammerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ScammerTag
	for rows.Next() {
		var i ScammerTag
		if err := rows.Scan(&i.ScammerID, &i.Tag); err != nil {
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

const updateCallHistory = `-- name: UpdateCallHistory :exec
UPDATE callhistory
SET
  scammer_id = $2,
  call_at = $3,
  call_time = $4,
  result = $5,
  talk_time = $6
WHERE id = $1
`

type UpdateCallHistoryParams struct {
	ID        string
	ScammerID string
	CallAt    time.Time
	CallTime  int64
	Result    bool
	TalkTime  sql.NullInt64
}

func (q *Queries) UpdateCallHistory(ctx context.Context, arg UpdateCallHistoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCallHistory,
		arg.ID,
		arg.ScammerID,
		arg.CallAt,
		arg.CallTime,
		arg.Result,
		arg.TalkTime,
	)
	return err
}

const updateMatching = `-- name: UpdateMatching :exec
UPDATE matching
SET
  created_at = $2,
  serial_number = $3,
  matched = $4,
  checked = $5,
  matching_at = $6,
  talk_time = $7,
  transcript = $8
WHERE id = $1
`

type UpdateMatchingParams struct {
	ID           string
	CreatedAt    time.Time
	SerialNumber int64
	Matched      bool
	Checked      bool
	MatchingAt   sql.NullTime
	TalkTime     sql.NullInt64
	Transcript   sql.NullString
}

func (q *Queries) UpdateMatching(ctx context.Context, arg UpdateMatchingParams) error {
	_, err := q.db.ExecContext(ctx, updateMatching,
		arg.ID,
		arg.CreatedAt,
		arg.SerialNumber,
		arg.Matched,
		arg.Checked,
		arg.MatchingAt,
		arg.TalkTime,
		arg.Transcript,
	)
	return err
}

const updateMetric = `-- name: UpdateMetric :exec
UPDATE metric
SET
  created_at = $2,
  calls = $3,
  scammers = $4,
  inactives = $5,
  call_time = $6,
  talk_time = $7,
  amount = $8
WHERE id = $1
`

type UpdateMetricParams struct {
	ID        string
	CreatedAt time.Time
	Calls     int64
	Scammers  int64
	Inactives int64
	CallTime  int64
	TalkTime  int64
	Amount    int64
}

func (q *Queries) UpdateMetric(ctx context.Context, arg UpdateMetricParams) error {
	_, err := q.db.ExecContext(ctx, updateMetric,
		arg.ID,
		arg.CreatedAt,
		arg.Calls,
		arg.Scammers,
		arg.Inactives,
		arg.CallTime,
		arg.TalkTime,
		arg.Amount,
	)
	return err
}

const updateScammer = `-- name: UpdateScammer :exec
UPDATE scammer
SET
  name = $2, 
  tel = $3, 
  is_active = $4
WHERE id = $1
`

type UpdateScammerParams struct {
	ID       string
	Name     string
	Tel      string
	IsActive bool
}

func (q *Queries) UpdateScammer(ctx context.Context, arg UpdateScammerParams) error {
	_, err := q.db.ExecContext(ctx, updateScammer,
		arg.ID,
		arg.Name,
		arg.Tel,
		arg.IsActive,
	)
	return err
}

const updateScammerTag = `-- name: UpdateScammerTag :exec
UPDATE scammer_tag
SET tag = $2
WHERE scammer_id = $1
`

type UpdateScammerTagParams struct {
	ScammerID string
	Tag       string
}

func (q *Queries) UpdateScammerTag(ctx context.Context, arg UpdateScammerTagParams) error {
	_, err := q.db.ExecContext(ctx, updateScammerTag, arg.ScammerID, arg.Tag)
	return err
}
