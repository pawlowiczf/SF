package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCountrySwiftCodeDetailsResponse struct {
	CountryISO2 string   `json:"countryISO2"`
	CountryName string   `json:"countryName"`
	SwiftCodes  []Branch `json:"swiftCodes"`
}

func (server *Server) GetCountrySwiftCodeDetails(ctx *gin.Context) {
	var response GetCountrySwiftCodeDetailsResponse

	countryISO2 := ctx.Param("country-iso2-code")

	err := validate.Var(countryISO2, "iso3166_1_alpha2")
	if err != nil {
		desc := "Provided ISO2Code is not valid"
		ctx.JSON(http.StatusBadRequest, errorResponseDesc(err, BadURIRequestArguments, desc))
		return 
	}

	swifts, err := server.store.GetCountrySwiftCodeDetails(ctx, countryISO2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, DatabaseInternalError))
		return
	}

	if len(swifts) == 0 {
		err := fmt.Errorf("no banks in %s country", countryISO2)
		ctx.JSON(http.StatusBadRequest, errorResponse(err, BadRequestArguments))
		return 
	}

	banks := parseBranches(swifts)
	response.CountryISO2 = countryISO2
	response.CountryName = swifts[0].CountryName
	response.SwiftCodes = banks

	ctx.JSON(http.StatusOK, response)
}
