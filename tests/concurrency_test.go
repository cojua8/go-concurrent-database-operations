package tests

import (
	"database/sql"
	"testing"

	"sandbox.com/concurrency/src/withmutex"
	"sandbox.com/concurrency/src/withtransactiondefault"
	"sandbox.com/concurrency/src/withtransactionisolation"
)

func TestDefaultTransaction(t *testing.T) {
	testCases := []struct {
		desc string
		work func(db *sql.DB)
	}{
		{
			desc: "execute with transaction defaults",
			work: withtransactiondefault.DoWork,
		},
		{
			desc: "execute with mutex",
			work: withmutex.DoWork,
		},
		{
			desc: "execute with transaction repeatable read isolation",
			work: withtransactionisolation.DoWork,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			db, err := InitDB()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			tC.work(db)

			// test for remaining stock
			var stock int
			if err := db.QueryRow("SELECT stock FROM products WHERE id = 1").Scan(&stock); err != nil {
				t.Fatal(err)
			}
			if stock != 0 {
				t.Errorf("expected stock to be 0, got %d", stock)
			}

			// test for created orders
			var orderCount int
			if err := db.QueryRow("SELECT COUNT(*) FROM orders WHERE product_id = 1").Scan(&orderCount); err != nil {
				t.Fatal(err)
			}
			if orderCount != 1 {
				t.Errorf("expected 1 order to be created, got %d", orderCount)
			}

		})
	}
}
