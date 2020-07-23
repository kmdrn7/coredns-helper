package models

import (
	"database/sql"
)

type Domain struct {
	ID int `json:"id"`
	Name string	`json:"name"`
	Master sql.NullString `json:"master"`
	LastCheck sql.NullInt64 `json:"last_check"`
	Type string `json:"type"`
	NotifiedSerial sql.NullInt64 `json:"notified_serial"`
	Account sql.NullString `json:"account"`
}

type DomainCollection struct {
	Domains []Domain `json:"data"`
}

func GetDomains(db *sql.DB) DomainCollection {
	sql := "SELECT * FROM domains"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := DomainCollection{}
	for rows.Next() {
		domain := Domain{}
		err2 := rows.Scan(&domain.ID, &domain.Name, &domain.Master, &domain.LastCheck, &domain.Type, &domain.NotifiedSerial, &domain.Account)
		if err2 != nil {
			panic(err2)
		}
		result.Domains = append(result.Domains, domain)
	}
	return result
}

func PutDomain(db *sql.DB, name string, type_domain string) (int64, error) {
	sql := "INSERT INTO domains(name, type) VALUES(?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(name, type_domain)
	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func EditDomain(db *sql.DB, id int, name string, type_domain string) (int64, error) {
	sql := "UPDATE domains set name = ?, type = ? WHERE id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err2 := stmt.Exec(name, type_domain, id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}

func DeleteDomain(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM domains WHERE id = ?"
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
