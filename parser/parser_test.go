package parser

import (
	"swift/swift"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCSV(t *testing.T) {
	parser := Parser{}

	swifts, err := parser.ParseCSV("../swift_codes.csv")
	require.NoError(t, err)
	require.NotEmpty(t, swifts)

	expectedA := swift.SwiftCSV{
		CountryISO2:   "AL",
		SwiftCode:     "AAISALTRXXX",
		BankName:      "UNITED BANK OF ALBANIA SH.A",
		Address:       "HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",
		CountryName:   "ALBANIA",
		IsHeadquarter: true,
	}
	expectedB := swift.SwiftCSV{
		CountryISO2:   "BG",
		SwiftCode:     "ABIEBGS1XXX",
		BankName:      "ABV INVESTMENTS LTD",
		Address:       "TSAR ASEN 20  VARNA, VARNA, 9002",
		CountryName:   "BULGARIA",
		IsHeadquarter: true,
	}

	actualA := swifts[0]
	actualB := swifts[1] 

	require.NotEmpty(t, actualA)
	require.NotEmpty(t, actualB)

	require.Equal(t, expectedA.Address, actualA.Address)
	require.Equal(t, expectedA.BankName, actualA.BankName)
	require.Equal(t, expectedA.CountryISO2, actualA.CountryISO2)
	require.Equal(t, expectedA.IsHeadquarter, actualA.IsHeadquarter)
	require.Equal(t, expectedA.SwiftCode, actualA.SwiftCode)
	require.Equal(t, expectedA.CountryName, actualA.CountryName)

	require.Equal(t, expectedB.Address, actualB.Address)
	require.Equal(t, expectedB.BankName, actualB.BankName)
	require.Equal(t, expectedB.CountryISO2, actualB.CountryISO2)
	require.Equal(t, expectedB.IsHeadquarter, actualB.IsHeadquarter)
	require.Equal(t, expectedB.SwiftCode, actualB.SwiftCode)
	require.Equal(t, expectedB.CountryName, actualB.CountryName)
}
