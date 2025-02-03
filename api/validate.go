package api

import (
	"fmt"
	"swift/parser"

	"github.com/gin-gonic/gin"
)

func validateHeadquarterSwiftCodeMismatch(swiftCode string, isHeadquarter bool) gin.H {
	if !parser.CheckIfHeadquarter(swiftCode) && isHeadquarter {
		err := fmt.Errorf("isHeadquarter nad swiftCode mismatch")
		errDesc := "Provided swiftCode is not associated with a headquarter (doesn't have XXX suffix)"
		return errorResponse(err, errDesc)
	}
	return nil 
}
