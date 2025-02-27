// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: swifts.sql

package db

import (
	"context"
)

const deleteSwiftCodeDetails = `-- name: DeleteSwiftCodeDetails :one
DELETE FROM swifts 
WHERE swift_code = $1
RETURNING swift_code, bank_name, country_iso2, country_name, address, is_headquarter
`

func (q *Queries) DeleteSwiftCodeDetails(ctx context.Context, swiftCode string) (Swift, error) {
	row := q.db.QueryRowContext(ctx, deleteSwiftCodeDetails, swiftCode)
	var i Swift
	err := row.Scan(
		&i.SwiftCode,
		&i.BankName,
		&i.CountryIso2,
		&i.CountryName,
		&i.Address,
		&i.IsHeadquarter,
	)
	return i, err
}

const getAllBranches = `-- name: GetAllBranches :many
SELECT swift_code, bank_name, country_iso2, country_name, address, is_headquarter FROM swifts
WHERE swift_code LIKE CONCAT($1::varchar, '%')
AND swift_code NOT LIKE CONCAT($1::varchar, '%XXX')
`

func (q *Queries) GetAllBranches(ctx context.Context, swiftcodeprefix string) ([]Swift, error) {
	rows, err := q.db.QueryContext(ctx, getAllBranches, swiftcodeprefix)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Swift{}
	for rows.Next() {
		var i Swift
		if err := rows.Scan(
			&i.SwiftCode,
			&i.BankName,
			&i.CountryIso2,
			&i.CountryName,
			&i.Address,
			&i.IsHeadquarter,
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

const getCountrySwiftCodeDetails = `-- name: GetCountrySwiftCodeDetails :many
SELECT swift_code, bank_name, country_iso2, country_name, address, is_headquarter FROM swifts 
WHERE country_iso2 = $1
`

func (q *Queries) GetCountrySwiftCodeDetails(ctx context.Context, countryIso2 string) ([]Swift, error) {
	rows, err := q.db.QueryContext(ctx, getCountrySwiftCodeDetails, countryIso2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Swift{}
	for rows.Next() {
		var i Swift
		if err := rows.Scan(
			&i.SwiftCode,
			&i.BankName,
			&i.CountryIso2,
			&i.CountryName,
			&i.Address,
			&i.IsHeadquarter,
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

const getRowsNumber = `-- name: GetRowsNumber :one
SELECT COUNT(*) FROM swifts
`

func (q *Queries) GetRowsNumber(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getRowsNumber)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getSwiftCodeDetails = `-- name: GetSwiftCodeDetails :one
SELECT swift_code, bank_name, country_iso2, country_name, address, is_headquarter FROM swifts 
WHERE swift_code = $1
LIMIT 1
`

func (q *Queries) GetSwiftCodeDetails(ctx context.Context, swiftCode string) (Swift, error) {
	row := q.db.QueryRowContext(ctx, getSwiftCodeDetails, swiftCode)
	var i Swift
	err := row.Scan(
		&i.SwiftCode,
		&i.BankName,
		&i.CountryIso2,
		&i.CountryName,
		&i.Address,
		&i.IsHeadquarter,
	)
	return i, err
}

const insertSwiftCodeDetails = `-- name: InsertSwiftCodeDetails :one
INSERT INTO swifts(swift_code, bank_name, country_iso2, country_name, address, is_headquarter)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING swift_code, bank_name, country_iso2, country_name, address, is_headquarter
`

type InsertSwiftCodeDetailsParams struct {
	SwiftCode     string `json:"swift_code"`
	BankName      string `json:"bank_name"`
	CountryIso2   string `json:"country_iso2"`
	CountryName   string `json:"country_name"`
	Address       string `json:"address"`
	IsHeadquarter bool   `json:"is_headquarter"`
}

func (q *Queries) InsertSwiftCodeDetails(ctx context.Context, arg InsertSwiftCodeDetailsParams) (Swift, error) {
	row := q.db.QueryRowContext(ctx, insertSwiftCodeDetails,
		arg.SwiftCode,
		arg.BankName,
		arg.CountryIso2,
		arg.CountryName,
		arg.Address,
		arg.IsHeadquarter,
	)
	var i Swift
	err := row.Scan(
		&i.SwiftCode,
		&i.BankName,
		&i.CountryIso2,
		&i.CountryName,
		&i.Address,
		&i.IsHeadquarter,
	)
	return i, err
}
