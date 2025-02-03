package api

import (
	"net/http"
	"strings"
	db "swift/db/sqlc"

	"github.com/gin-gonic/gin"
)

type InsertSwiftCodeDetailsRequest struct {
	Address       string `json:"address"       binding:"required"`
	BankName      string `json:"bankName"      binding:"required"`
	CountryISO2   string `json:"countryISO2"   binding:"required,iso3166_1_alpha2"`
	CountryName   string `json:"countryName"   binding:"required"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	SwiftCode     string `json:"swiftCode"     binding:"required,bic"`
}

func (server *Server) InsertSwiftCodeDetails(ctx *gin.Context) {
	var req InsertSwiftCodeDetailsRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, BadRequestArguments))
		return
	}

	if errResponse := validateHeadquarterSwiftCodeMismatch(req.SwiftCode, req.IsHeadquarter); errResponse != nil {
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	req.BankName = strings.ToUpper(req.BankName)
	req.CountryISO2 = strings.ToUpper(req.CountryISO2)
	req.Address = strings.ToUpper(req.Address)
	req.CountryName = strings.ToUpper(req.CountryName)

	_, err = server.store.InsertSwiftCodeDetails(ctx, db.InsertSwiftCodeDetailsParams{
		SwiftCode:     req.SwiftCode,
		BankName:      req.BankName,
		CountryIso2:   req.CountryISO2,
		CountryName:   req.CountryName,
		Address:       req.Address,
		IsHeadquarter: req.IsHeadquarter,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, DatabaseInternalError))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully inserted new SWIFT code entry to the database",
	})
}
