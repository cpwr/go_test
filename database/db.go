package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

var db *sql.DB

func CreateConn() (*sql.DB, error) {

    connStr := "postgres://gotest:p@ssw0rd@localhost:5432/gotest?sslmode=disable"
    db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func CreateTable() error {

    _, err := db.Exec("CREATE TABLE employees (name VARCHAR(20), created DATETIME)");
    if err != nil {
        return err
    }

    _, err = db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`);
    if err != nil {
        return err
    }

    return nil
}
