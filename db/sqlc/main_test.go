package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	DB_DRIVER = "postgres"
	DB_SOURCE = "postgresql://root:secret@localhost:5432/account?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	conn, err := sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatalln(err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
