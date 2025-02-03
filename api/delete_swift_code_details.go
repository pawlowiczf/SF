package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) DeleteSwiftCodeDetails(ctx *gin.Context) {
	swiftCode := ctx.Param("swift-code")

	err := validate.Var(swiftCode, "bic,gte=8,lte=11")
	if err != nil {
		desc := "Provided SwiftCode doesn't exist"
		ctx.JSON(http.StatusBadRequest, errorResponseDesc(err, BadURIRequestArguments, desc))
		return
	}

	_, err = server.store.DeleteSwiftCodeDetails(ctx, swiftCode)
	if err != nil {
		if err == sql.ErrNoRows {
			desc := "The entry, with provided swiftCode, doesn't exists in the database" 
			ctx.JSON(http.StatusBadRequest, errorResponseDesc(err, DatabaseNotFoundRows, desc))
			return 
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, DatabaseInternalError))
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted swiftCode details",
	})
}
