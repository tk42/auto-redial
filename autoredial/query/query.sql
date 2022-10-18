-- name: ListMetric :many
SELECT * FROM metric
WHERE created_at BETWEEN $1 AND $2
ORDER BY created_at DESC;

-- name: CreateMetric :one
INSERT INTO metric (
  id, created_at, calls, scammers, inactives, call_sec, talk_sec, amount
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateMetric :exec
UPDATE metric
SET
  created_at = $2,
  calls = $3,
  scammers = $4,
  inactives = $5,
  call_sec = $6,
  talk_sec = $7,
  amount = $8
WHERE id = $1;

-- name: DeleteMetric :exec
DELETE FROM metric
WHERE created_at BETWEEN $1 AND $2;


-- name: GetScammer :one
SELECT * FROM scammer
WHERE id = $1;

-- name: CreateScammer :one
INSERT INTO scammer (
  id, name, tel, is_active
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateScammer :exec
UPDATE scammer
SET
  name = $2, 
  tel = $3, 
  is_active = $4
WHERE id = $1;

-- name: DeleteScammer :exec
DELETE FROM scammer
WHERE id = $1;


-- name: GetCallHistory :one
SELECT * FROM callhistory
WHERE id = $1;

-- name: ListCallHistoryByScammerId :many
SELECT * FROM callhistory
WHERE scammer_id = $1;

-- name: CreateCallHistory :one
INSERT INTO callhistory (
  id, scammer_id, call_at, call_sec, result, talk_sec
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateCallHistory :exec
UPDATE callhistory
SET
  scammer_id = $2,
  call_at = $3,
  call_sec = $4,
  result = $5,
  talk_sec = $6
WHERE id = $1;

-- name: DeleteCallHistory :exec
DELETE FROM callhistory
WHERE id = $1;



-- name: GetMatching :one
SELECT * FROM matching
WHERE id = $1;

-- name: ListMatching :many
SELECT * FROM matching
WHERE checked = matched;

-- name: CreateMatching :one
INSERT INTO matching (
  id, created_at, serial_number, matched, checked, matching_at, talk_sec, transcript
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateMatchingMatched :exec
UPDATE matching
SET
  matched = $2,
  matching_at = $3,
  talk_sec = $4
WHERE id = $1;

-- name: UpdateMatchingChecked :exec
UPDATE matching
SET
  checked = $2
WHERE id = $1;

-- name: UpdateMatchingTranscript :exec
UPDATE matching
SET
  transcript = $2
WHERE id = $1;

-- name: DeleteMatching :exec
DELETE FROM matching
WHERE id = $1;


-- name: ListScammerTag :many
SELECT * FROM scammer_tag
WHERE scammer_id = $1;

-- name: CreateScammerTag :one
INSERT INTO scammer_tag (
  scammer_id, tag
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateScammerTag :exec
UPDATE scammer_tag
SET tag = $2
WHERE scammer_id = $1;

-- name: DeleteScammerTag :exec
DELETE FROM scammer_tag
WHERE scammer_id = $1 AND tag = $2;


-- name: ListScammerMatchingByScammerId :many
SELECT * FROM scammer_matching
WHERE scammer_id = $1;

-- name: ListScammerMatchingByMatchingId :many
SELECT * FROM scammer_matching
WHERE matching_id = $1;

-- name: CreateScammerMatching :one
INSERT INTO scammer_matching (
  scammer_id, matching_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteScammerMatching :exec
DELETE FROM scammer_matching
WHERE scammer_id = $1 AND matching_id = $2;


-- name: ListScammerCallByScammerId :many
SELECT * FROM scammer_call
WHERE scammer_id = $1;

-- name: ListScammerCallByCallId :many
SELECT * FROM scammer_call
WHERE call_id = $1;

-- name: CreateScammerCall :one
INSERT INTO scammer_call (
  scammer_id, call_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteScammerCall :exec
DELETE FROM scammer_call
WHERE scammer_id = $1 AND call_id = $2;


-- name: ListMatchingTag :many
SELECT * FROM matching_tag
WHERE matching_id = $1;