package tests

import (
	"database/sql"
	"fmt"
	"os"

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
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s",
		os.Getenv("POSTGRES_HOSTNAME"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"))
	db, err := sql.Open("pgx", connectionString)
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
