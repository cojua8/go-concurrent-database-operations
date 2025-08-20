package tests

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() (*sql.DB, error) {
	db, err := connectDb()
	if err != nil {
		return nil, err
	}

	err = resetDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDb() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=db user=postgres dbname=postgres password=postgres")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func resetDB(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET stock = 2 ")
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM orders")
	if err != nil {
		return err
	}

	return nil
}
