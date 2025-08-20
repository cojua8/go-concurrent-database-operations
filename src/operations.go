package src

import (
	"database/sql"
	"fmt"
	"time"
)

type Query interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
}

func Process(q Query, who int, waitBeforeStock, waitBeforeOrders, waitBeforeUpdate time.Duration) error {
	time.Sleep(waitBeforeStock)
	stock, err := GetProductStock(q)
	if err != nil {
		return err
	}
	fmt.Println("User", who, "stock:", stock)

	if stock < 2 {
		return fmt.Errorf("not enough stock")
	}

	time.Sleep(waitBeforeOrders)
	if err := InsertOrder(q); err != nil {
		return err
	}
	fmt.Println("User", who, "inserted order")

	time.Sleep(waitBeforeUpdate)
	if err := UpdateProductStock(q); err != nil {
		return err
	}
	fmt.Println("User", who, "updated stock")

	return nil
}

func GetProductStock(q Query) (int, error) {
	var stock int
	if err := q.QueryRow("SELECT stock FROM products WHERE id = 1").Scan(&stock); err != nil {
		return 0, err
	}
	return stock, nil
}

func InsertOrder(q Query) error {
	if _, err := q.Exec("INSERT INTO orders (product_id, amount) VALUES (1, 2)"); err != nil {
		return err
	}
	return nil
}

func UpdateProductStock(q Query) error {
	if _, err := q.Exec("UPDATE products SET stock = stock - 2 WHERE id = 1"); err != nil {
		return err
	}
	return nil
}
