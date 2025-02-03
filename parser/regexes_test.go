package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckIfHeadquarter(t *testing.T) {
	codes := []string{
		"AAISALTRXXX",
		"ABIEBGS1XXX",
		"ADCRBGS1XXX",
		"AKBKMTMTXXX",
		"XXX",
	}
	for _, code := range codes {
		flag := CheckIfHeadquarter(code)
		require.Equal(t, true, flag)
	}

	codes = []string{
		"ALBPPLP1BMW",
		"XXXALBPPLP1BMW",
		"BCHICLR10R2",
		"BCHICLRMEXP",
		"XXXA",
	}
	for _, code := range codes {
		flag := CheckIfHeadquarter(code)
		require.Equal(t, false, flag)
	}
}
