package main

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func main() {
	path := os.Getenv("APP_DB_PATH")
	if path == "" { path = "data/service.sqlite3" }
	_ = os.MkdirAll(filepath.Dir(path), 0o755)
	database, err := sql.Open("sqlite", path)
	if err != nil { panic(err) }
	defer database.Close()
	if _, err = database.Exec("CREATE TABLE IF NOT EXISTS schema_versions (version INTEGER PRIMARY KEY, applied_at TEXT NOT NULL)"); err != nil { panic(err) }
}
