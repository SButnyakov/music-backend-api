package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// Открывает бд для тестов и пингует
// Нужна, чтобы вызвать её через defer и закрыть соединение после тестов
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("sqlserver", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("DELETE %s", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}
