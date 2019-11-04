package database

import "database/sql"

type Database struct {
	db *sql.DB
}

func Connect(driverName, connectionString string) (*Database, error) {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Database{
		db: db,
	}, nil
}
