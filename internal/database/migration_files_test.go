package database

import (
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestMigrationFilesUseSortableThreeDigitPrefix(t *testing.T) {
	files, err := loadSQLFiles(filepath.Join("..", "..", "migrations"))
	if err != nil {
		t.Fatal(err)
	}

	pattern := regexp.MustCompile(`^\d{3}_[a-z0-9_]+\.sql$`)
	seenPrefixes := map[string]string{}
	for _, file := range files {
		name := filepath.Base(file)
		if !pattern.MatchString(name) {
			t.Fatalf("migration %q must use sortable three-digit prefix, e.g. 020_spawn_zone.sql", name)
		}
		prefix := strings.SplitN(name, "_", 2)[0]
		if previous, ok := seenPrefixes[prefix]; ok {
			t.Fatalf("migration prefix %s is duplicated by %q and %q", prefix, previous, name)
		}
		seenPrefixes[prefix] = name
	}
}
