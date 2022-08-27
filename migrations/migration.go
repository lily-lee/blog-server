package migrations

import (
	"sort"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations []*gormigrate.Migration

func Migrate(db *gorm.DB) error {
	return getGomigrate(db).Migrate()
}

func RollbackLast(db *gorm.DB) error {
	return getGomigrate(db).RollbackLast()
}

func getGomigrate(db *gorm.DB) *gormigrate.Gormigrate {
	sort.Slice(Migrations, func(i, j int) bool {
		n := Migrations
		return n[i].ID < n[j].ID
	})
	gormigrate.DefaultOptions.IDColumnSize = 50
	return gormigrate.New(
		db, /*.LogMode(true)*/
		gormigrate.DefaultOptions,
		Migrations,
	)
}
