package swift

// Codes ending with 'XXX' represent a bank's headquarters, otherwise branch
// Branch codes are associated with a headquarters if their first 8 characters match
// Codes can represent both the branch and thte headquarter of the bank

type SwiftCSV struct {
	CountryISO2   string `json:"countryISO2" validate:"alpha,gte=2,lte=2,iso3166_1_alpha2"`
	SwiftCode     string `json:"swiftCode" validate:"gte=8,lte=11,bic"`
	BankName      string `json:"bankName"`
	Address       string `json:"address"`
	CountryName   string `json:"countryName"`
	IsHeadquarter bool   `json:"isHeadquarter"`
}

// Country codes and names must always must always be stored and returned as uppercase strings
// Redundant columns in the file may be omitted
