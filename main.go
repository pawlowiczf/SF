package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"swift/api"
	"swift/config"
	db "swift/db/sqlc"
	"swift/parser"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config %v", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot open db %v", err)
	}
	if err = conn.Ping(); err != nil {
		log.Fatalf("cannot connect to db %v", err)
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	loadSwiftDetailsToDB(store, config.SwiftCSVPath)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create new HTTP server %v", err)
	}

	fmt.Println("Running HTTP server on", config.HTTPServerAddress)

	err = server.Run()
	if err != nil {
		log.Fatalf("cannot run HTTP server %v", err)
	}
}

func loadSwiftDetailsToDB(store db.Store, filename string) {
	amount, err := store.GetRowsNumber(context.Background())
	if err != nil {
		log.Fatalf("cannot query db %v", err)
	}
	if amount > 0 {
		return
	}

	parser := parser.Parser{}
	swiftsCSV, err := parser.ParseCSV(filename)
	if err != nil {
		log.Fatalf("cannot parse %s %v", filename, err)
	}

	for _, swiftCSV := range swiftsCSV {
		args := db.InsertSwiftCodeDetailsParams{
			SwiftCode:     swiftCSV.SwiftCode,
			BankName:      swiftCSV.BankName,
			CountryIso2:   swiftCSV.CountryISO2,
			CountryName:   swiftCSV.CountryName,
			Address:       swiftCSV.Address,
			IsHeadquarter: swiftCSV.IsHeadquarter,
		}

		_, err = store.InsertSwiftCodeDetails(context.Background(), args)
		if err != nil {
			log.Fatalf("cannot query db %v", err)
		}
	}

}

func runDBMigration(migrationURL string, dbSource string) {
	fmt.Println("Running DB migrations")
	
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatalf("cannot run db migrations %v", err)
	}

	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("cannot run up db migrations %v", err)
	}
}
