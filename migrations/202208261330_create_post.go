package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "202208261330_create_post",
		Migrate: func(db *gorm.DB) error {
			type Post struct {
				ID        uint64    `gorm:"primary_key;not null;comment:id"`
				UserID    uint64    `gorm:"index;not null;default:0;comment:user.id"`
				DraftID   uint64    `gorm:"index;not null;default:0;comment:draft.id"`
				VolumeID  uint64    `gorm:"index;not null;default:0;comment:volume.id"`
				Title     string    `gorm:"type:varchar(255);index:fulltext_title_digest_content,class:FULLTEXT,option:WITH PARSER ngram;not null;default:'';comment:title"`
				Content   string    `gorm:"type:longtext;index:fulltext_title_digest_content,class:FULLTEXT,option:WITH PARSER ngram;not null;comment:content"`
				Digest    string    `gorm:"type:varchar(300);index:fulltext_title_digest_content,class:FULLTEXT,option:WITH PARSER ngram;not null;default:'';comment:digest"`
				CoverURL  string    `gorm:"type:varchar(200);not null;default:'';comment:cover url"`
				Tag       string    `gorm:"type:varchar(255);not null;default:'';comment:tag"`
				CreatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);comment:create time"`
				UpdatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:update time"`
			}

			return db.Migrator().AutoMigrate(&Post{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("posts")
		},
	})
}
