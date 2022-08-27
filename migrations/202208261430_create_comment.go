package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "202208261430_create_comment",
		Migrate: func(db *gorm.DB) error {
			type Comment struct {
				ID        uint64    `gorm:"primary_key;not null;comment:id"`
				PostID    uint64    `gorm:"index:idx_post_id_pid;not null;default:0;comment:post.id"`
				PID       uint64    `gorm:"index:idx_post_id_pid;not null;default:0;comment:parent comment id"`
				UserID    uint64    `gorm:"index;not null;default:0;comment:user.id"`
				Content   string    `gorm:"type:varchar(1024);not null;default:'';comment:content"`
				Anonymous bool      `gorm:"type:tinyint(1);not null;default:0;comment:anonymous"`
				CreatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);comment:create time"`
				UpdatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:update time"`
			}

			return db.Migrator().AutoMigrate(&Comment{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("comments")
		},
	})
}
