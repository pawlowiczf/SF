package parser

import (
	"strings"
)

// var CheckIfHeadquarter = regexp.MustCompile(`XXX$`).MatchString

var CheckIfHeadquarter = func(swiftCode string) bool {
	return strings.HasSuffix(swiftCode, "XXX")
}
