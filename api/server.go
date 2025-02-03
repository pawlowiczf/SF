package api

import (
	"swift/config"
	db "swift/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Server struct {
	config config.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	server := &Server{}
	server.config = config
	server.store = store

	server.setupServerRoutes()

	return server, nil
}

func (server *Server) setupServerRoutes() {
	router := gin.Default()

	router.GET("/v1/swift-codes/:swift-code", server.GetSwiftCodeDetails)
	router.GET("/v1/swift-codes/country/:country-iso2-code", server.GetCountrySwiftCodeDetails)
	router.POST("/v1/swift-codes/", server.InsertSwiftCodeDetails)
	router.DELETE("/v1/swift-codes/:swift-code", server.DeleteSwiftCodeDetails)

	server.router = router
}

func (server *Server) Run() error {
	return server.router.Run(server.config.HTTPServerAddress)
}
