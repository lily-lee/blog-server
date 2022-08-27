package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "202208261400_create_post_like",
		Migrate: func(db *gorm.DB) error {
			type PostLike struct {
				PostID    uint64    `gorm:"index:idx_post_id_user_id;not null;default:0;comment:post.id"`
				UserID    uint64    `gorm:"index:idx_post_id_user_id;not null;default:0;comment:user.id"`
				CreatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);comment:create time"`
				UpdatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:update time"`
			}

			return db.Migrator().AutoMigrate(&PostLike{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("post_likes")
		},
	})
}
