package utils

import (
	"database/sql"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	migrate(db)
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS domains(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(255) NOT NULL,
				master VARCHAR(128),
				last_check BIGINT,
				type VARCHAR(6) NOT NULL,
				notified_serial BIGINT,
				account VARCHAR(40)
    );
		CREATE TABLE IF NOT EXISTS records(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				domain_id BIGINT,
				name VARCHAR(255),
				type VARCHAR(10),
				content VARCHAR(600),
				ttl INTEGER,
				prio INTEGER,
				chang_data INTEGER,
				disabled BOOL
		);
	`
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}