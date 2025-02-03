package parser

import (
	"encoding/csv"
	"os"
	"strings"
	"swift/swift"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Parser struct {
}

func (parser *Parser) ParseCSV(filename string) ([]swift.SwiftCSV, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	swifts := make([]swift.SwiftCSV, len(records))
	for idx, record := range records[1:] {
		swifts[idx] = swift.SwiftCSV{
			CountryISO2:   record[0],
			SwiftCode:     record[1],
			BankName:      strings.ToUpper(record[3]),
			Address:       strings.ToUpper(record[4]),
			CountryName:   strings.ToUpper(record[6]),
			IsHeadquarter: CheckIfHeadquarter(record[1]),
		}
		if swifts[idx].Address == "  " {
			swifts[idx].Address = ""
		}

		err := validate.Struct(swifts[idx])
		if err != nil {
			return nil, err
		}
	}

	return swifts, nil
}
