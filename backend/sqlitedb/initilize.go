package sqlitedb

import (
	"database/sql"
	"os"
	"path/filepath"
)

// Reads a SQL file and executes it on the given DB connection
func executeSLQFile(db *sql.DB, path string) error {
	
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return err
	}
	return nil
}

// Reads and executes all schema files
func InitializeSchemas(db *sql.DB, schemaDir string) error {
	files, err := os.ReadDir(schemaDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			
			err := executeSLQFile(db, filepath.Join(schemaDir, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}