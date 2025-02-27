// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	DeleteSwiftCodeDetails(ctx context.Context, swiftCode string) (Swift, error)
	GetAllBranches(ctx context.Context, swiftcodeprefix string) ([]Swift, error)
	GetCountrySwiftCodeDetails(ctx context.Context, countryIso2 string) ([]Swift, error)
	GetRowsNumber(ctx context.Context) (int64, error)
	GetSwiftCodeDetails(ctx context.Context, swiftCode string) (Swift, error)
	InsertSwiftCodeDetails(ctx context.Context, arg InsertSwiftCodeDetailsParams) (Swift, error)
}

var _ Querier = (*Queries)(nil)
