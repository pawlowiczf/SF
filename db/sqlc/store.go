package db

import "database/sql"

type Store interface {
	Querier
}

func NewStore(db *sql.DB) Store {
	return SQLStore{
		Queries: New(db),
		db:      db,
	}
}

type SQLStore struct {
	*Queries 
	db *sql.DB
}
