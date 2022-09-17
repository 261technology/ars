package db

import (
	"log"

	"github.com/harisaginting/guin/db/table"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	log.Println("migration db ", db.Migrator().CurrentDatabase())
	table.MigrateShorten(db)
	log.Println("migration success")
}
