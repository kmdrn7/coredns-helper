package models

import (
	"database/sql"
)

type Record struct {
	ID int `json:"id"`
	DomainID int `json:"domain_id"`
	Name string	`json:"name"`
	Type string `json:"type"`
	Content string `json:"content"`
	Ttl int `json:"ttl"`
	Prio sql.NullInt32 `json:"prio"`
	ChangDate sql.NullInt32 `json:"chang_date"`
	Disabled bool `json:"disabled"`
}

type RecordCollection struct {
	Records []Record `json:"data"`
}

func GetRecords(db *sql.DB) RecordCollection {
	sql := "SELECT * FROM records"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := RecordCollection{}
	for rows.Next() {
		record := Record{}
		err2 := rows.Scan(&record.ID, &record.DomainID, &record.Name, &record.Type, &record.Content, &record.Ttl, &record.Prio, &record.ChangDate, &record.Disabled)
		if err2 != nil {
			panic(err2)
		}
		result.Records = append(result.Records, record)
	}
	return result
}

func PutRecords(db *sql.DB, domain_id int, name string, type_record string, content string, ttl int, disabled bool) (int64, error) {
	sql := "INSERT INTO records (domain_id, name, type, content, ttl, disabled) VALUES(?,?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(domain_id, name, type_record, content, ttl, disabled)
	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func EditRecords(db *sql.DB, id int, name string, type_record string, content string, ttl int, disabled bool) (int64, error) {
	sql := "UPDATE records set name = ?, type = ?, content = ?, ttl = ?, disabled = ? WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err2 := stmt.Exec(name, type_record, content, ttl, disabled, id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}

func DeleteRecords(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM records WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
