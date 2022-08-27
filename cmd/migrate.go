package cmd

import (
	"github.com/lily-lee/blog-server/config"
	"github.com/lily-lee/blog-server/migrations"
)

func Migrate() error {
	return migrations.Migrate(config.DB)
}

func Rollback() error {
	return migrations.RollbackLast(config.DB)
}
