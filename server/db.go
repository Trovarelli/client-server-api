package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./cotacoes.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		name TEXT,
		high REAL,
		low REAL,
		varBid REAL,
		pctChange REAL,
		bid REAL,
		ask REAL,
		timestamp INTEGER,
		created_at DATETIME
	);`

	_, err = db.Exec(query)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
