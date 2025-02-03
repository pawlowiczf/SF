package api

import "github.com/gin-gonic/gin"

const (
	BadRequestArguments    = "Provided bad request arguments"
	BadURIRequestArguments = "Provided bad request URI arguments"
	DatabaseInternalError  = "Error with database query/connection"
	DatabaseNotFoundRows   = "The query resulted in no rows set"
)

func errorResponseDesc(err error, errorDescription string, description string) gin.H {
	return gin.H{
		"ErrorDescription": errorDescription,
		"Description":      description,
		"Error":            err.Error(),
	}
}

func errorResponse(err error, errorDescription string) gin.H {
	return gin.H{
		"ErrorDescription": errorDescription,
		"Error":            err.Error(),
	}
}
