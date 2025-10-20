package data

import (
	"leo/internal/data/model"
	"os"
	"path/filepath"

	"github.com/google/wire"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData() (*Data, func(), error) {
	dir, _ := os.UserHomeDir()
	dbPath := filepath.Join(dir, ".peon/leo.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	d := &Data{db: db}
	cleanup := func() {
		if db, err := d.db.DB(); err == nil {
			_ = db.Close()
		}
	}

	tables := []any{
		model.Snippet{},
	}

	for _, t := range tables {
		if err := db.AutoMigrate(&t); err != nil {
			return nil, cleanup, err
		}
	}

	return d, cleanup, nil
}
