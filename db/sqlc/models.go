// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import ()

type Swift struct {
	SwiftCode     string `json:"swift_code"`
	BankName      string `json:"bank_name"`
	CountryIso2   string `json:"country_iso2"`
	CountryName   string `json:"country_name"`
	Address       string `json:"address"`
	IsHeadquarter bool   `json:"is_headquarter"`
}
