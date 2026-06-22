package database

import (
	"path/filepath"
	"regexp"
	"testing"
)

func TestMigrationFilesUseSortableThreeDigitPrefix(t *testing.T) {
	files, err := loadSQLFiles(filepath.Join("..", "..", "migrations"))
	if err != nil {
		t.Fatal(err)
	}

	pattern := regexp.MustCompile(`^\d{3}_[a-z0-9_]+\.sql$`)
	for _, file := range files {
		name := filepath.Base(file)
		if !pattern.MatchString(name) {
			t.Fatalf("migration %q must use sortable three-digit prefix, e.g. 020_spawn_zone.sql", name)
		}
	}
}
