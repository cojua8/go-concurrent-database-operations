// Efectivo, pero depende de la infrastructura para que sea efectivo realmente
package withmutex

import (
	"database/sql"
	"fmt"
	"sync"

	"sandbox.com/concurrency/src"
)

func DoWork(db *sql.DB) {
	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := SellProductTransaction(mu, db, 1)
		if err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		defer wg.Done()
		err := SellProductTransaction(mu, db, 2)
		if err != nil {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}

func SellProductTransaction(mu *sync.Mutex, db *sql.DB, goroID int) error {
	mu.Lock()
	defer mu.Unlock()

	return src.Process(db, goroID, 0, 0, 0)
}
