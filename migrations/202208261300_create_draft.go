package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "202208261300_create_draft",
		Migrate: func(db *gorm.DB) error {
			type Draft struct {
				ID        uint64    `gorm:"primary_key;not null;comment:id"`
				UserID    uint64    `gorm:"index;not null;default:0;comment:users.id"`
				PostID    uint64    `gorm:"index;not null;default:0;comment:posts.id"`
				VolumeID  uint64    `gorm:"index;not null;default:0;comment:volumes.id"`
				Title     string    `gorm:"type:varchar(255);not null;default:'';comment:title"`
				Content   string    `gorm:"type:longtext;not null;comment:content"`
				Digest    string    `gorm:"type:varchar(300);not null;default:'';comment:digest"`
				CoverURL  string    `gorm:"type:varchar(200);not null;default:'';comment:cover url"`
				Tag       string    `gorm:"type:varchar(255);not null;default:'';comment:tag"`
				Posted    bool      `gorm:"type:tinyint(1);not null;default:0;comment:post status"`
				CreatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);comment:create time"`
				UpdatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:update time"`
			}

			return db.Migrator().AutoMigrate(&Draft{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("drafts")
		},
	})
}
