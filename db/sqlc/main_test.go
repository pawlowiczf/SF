package db

import (
	"database/sql"
	"log"
	"os"
	"swift/config"
	"testing"

	_ "github.com/lib/pq"
)

var testDB *sql.DB
var store Store

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../.")
	if err != nil {
		log.Fatal("cannot load config during tests: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db during tests: ", err)
	}
	if err = testDB.Ping(); err != nil {
		log.Fatal("cannot ping/connect to db during tests: ", err)
	}

	store = NewStore(testDB)
	os.Exit(m.Run())
}
