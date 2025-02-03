-- name: GetSwiftCodeDetails :one 
SELECT * FROM swifts 
WHERE swift_code = $1
LIMIT 1;

-- name: GetRowsNumber :one 
SELECT COUNT(*) FROM swifts;

-- name: InsertSwiftCodeDetails :one 
INSERT INTO swifts(swift_code, bank_name, country_iso2, country_name, address, is_headquarter)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllBranches :many 
SELECT * FROM swifts
WHERE swift_code LIKE CONCAT(sqlc.arg(swiftCodePrefix)::varchar, '%')
AND swift_code NOT LIKE CONCAT(sqlc.arg(swiftCodePrefix)::varchar, '%XXX');

-- name: GetCountrySwiftCodeDetails :many 
SELECT * FROM swifts 
WHERE country_iso2 = $1; 

-- name: DeleteSwiftCodeDetails :one 
DELETE FROM swifts 
WHERE swift_code = $1
RETURNING *;