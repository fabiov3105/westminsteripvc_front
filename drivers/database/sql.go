package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// banco de dados
func ConectDB() *sql.DB {
	stringDeConexao := "host=localhost user=admin password=admin dbname=westiminster_db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", stringDeConexao)
	if err != nil {
		panic(err.Error)
	}
	return db
}
