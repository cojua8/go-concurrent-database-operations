.PHONY: test

cleancache:
	go clean -testcache


testtransactiondefault: cleancache
	go test -v -timeout 30s -run ^TestDefaultTransaction/execute_with_transaction_defaults sandbox.com/concurrency/tests

testtransactionlock: cleancache
	go test -v -timeout 30s -run ^TestDefaultTransaction/execute_with_mutex sandbox.com/concurrency/tests

testwithmutex: cleancache
	go test -v -timeout 30s -run ^TestDefaultTransaction/execute_with_transaction_repeatable_read_isolation sandbox.com/concurrency/tests

runmigrations:
	migrate -database  "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOSTNAME}:5432/${POSTGRES_DB}?sslmode=disable" -path db/migrations/ up

deletemigrations:
	migrate -database  "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOSTNAME}:5432/${POSTGRES_DB}?sslmode=disable" -path db/migrations/ down