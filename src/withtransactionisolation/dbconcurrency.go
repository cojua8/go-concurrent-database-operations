// SPOILER: no es efectivo contra operaciones concurrentes
package withtransactionisolation

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"sandbox.com/concurrency/src"
)

func DoWork(db *sql.DB) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := SellProductTransaction(db, 50*time.Millisecond, 100*time.Millisecond, 100*time.Millisecond, 1)
		if err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		defer wg.Done()
		err := SellProductTransaction(db, 100*time.Millisecond, 100*time.Millisecond, 100*time.Millisecond, 2)
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}

func SellProductTransaction(db *sql.DB, waitBeforeStock, waitBeforeOrders, waitBeforeUpdate time.Duration, goroID int) error {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = src.Process(tx, goroID, waitBeforeStock, waitBeforeOrders, waitBeforeUpdate)
	if err != nil {
		return err
	}

	return tx.Commit()
}
