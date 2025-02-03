package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCountrySwiftCodeDetails(t *testing.T) {
	countries := map[string]string{"PL": "POLAND", "LV": "LATVIA", "CL": "CHILE"}

	for countryISO2, countryName := range countries {
		swifts, err := store.GetCountrySwiftCodeDetails(context.Background(), countryISO2)
		require.NoError(t, err)
		require.NotEmpty(t, swifts)

		for _, swift := range swifts {
			require.Equal(t,  countryISO2, swift.CountryIso2)
			require.Equal(t, countryName, swift.CountryName)
		}
	}
}

func TestGetSwiftCodeDetails(t *testing.T) {
	swiftsCodes := []string{"ALBPPLPWXXX", "BCECCLRMXXX", "BKSACLRMXXX"}

	for _, swiftCode := range swiftsCodes {
		swift, err := store.GetSwiftCodeDetails(context.Background(), swiftCode)
		require.NoError(t, err)
		require.NotEmpty(t, swift)
		require.Equal(t, swiftCode, swift.SwiftCode)
	}
}

