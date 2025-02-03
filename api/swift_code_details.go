package api

import (
	"database/sql"
	"net/http"
	db "swift/db/sqlc"

	"github.com/gin-gonic/gin"
)

type GetSwiftCodeDetailsResponse struct {
	Address       string   `json:"address"`
	BankName      string   `json:"bankName"`
	CountryISO2   string   `json:"countryISO2"`
	CountryName   string   `json:"countryName"`
	IsHeadquarter bool     `json:"isHeadquarter"`
	SwiftCode     string   `json:"swiftCode"`
	Branches      []Branch `json:"branches,omitempty"`
}

type Branch struct {
	Address       string `json:"address"`
	BankName      string `json:"bankName"`
	CountryISO2   string `json:"countryISO2"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	SwiftCode     string `json:"swiftCode"`
}

func (server *Server) GetSwiftCodeDetails(ctx *gin.Context) {
	var response GetSwiftCodeDetailsResponse
	response.Branches = nil

	swiftCode := ctx.Param("swift-code")

	err := validate.Var(swiftCode, "bic,gte=8,lte=11")
	if err != nil {
		desc := "Provided SwiftCode doesn't exist"
		ctx.JSON(http.StatusBadRequest, errorResponseDesc(err, BadURIRequestArguments, desc))
		return
	}

	swiftDB, err := server.store.GetSwiftCodeDetails(ctx, swiftCode)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponseDesc(err, DatabaseNotFoundRows, "No banks with such a SwiftCode"))
			return 
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, DatabaseInternalError))
		return
	}

	response = parseResponse(swiftDB)

	if swiftDB.IsHeadquarter {
		branchesDB, err := server.store.GetAllBranches(ctx, swiftDB.SwiftCode[:len(swiftDB.SwiftCode)-3])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err, DatabaseInternalError))
			return
		}
		response.Branches = parseBranches(branchesDB)
	}

	ctx.JSON(http.StatusOK, response)
}

func parseBranches(branchesDB []db.Swift) []Branch {
	var branches []Branch

	for _, branch := range branchesDB {
		branches = append(branches, Branch{
			Address:       branch.Address,
			BankName:      branch.BankName,
			CountryISO2:   branch.CountryIso2,
			IsHeadquarter: branch.IsHeadquarter,
			SwiftCode:     branch.SwiftCode,
		})
	}

	return branches
}

func parseResponse(swiftDB db.Swift) GetSwiftCodeDetailsResponse {
	return GetSwiftCodeDetailsResponse{
		Address:       swiftDB.Address,
		BankName:      swiftDB.BankName,
		CountryISO2:   swiftDB.CountryIso2,
		CountryName:   swiftDB.CountryName,
		IsHeadquarter: swiftDB.IsHeadquarter,
		SwiftCode:     swiftDB.SwiftCode,
	}
}
