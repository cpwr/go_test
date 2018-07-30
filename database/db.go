package database

import (
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)


func CreateConn(db *sql.DB) (*sql.DB, error) {
    var err error
    connStr := "postgres://gotest:p@ssw0rd@localhost:5432/gotest?sslmode=disable"
    db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB) error {

    _, err := db.Exec("CREATE TABLE IF NOT EXISTS employees (name VARCHAR(20), created TIMESTAMP)");
    if err != nil {
        return err
    }

    _, err = db.Exec(`INSERT INTO employees (name, created) values ('Aaron', NOW())`);
    if err != nil {
        return err
    }

    return nil
}
