package migrations

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "202208252000_create_user",
		Migrate: func(db *gorm.DB) error {
			type User struct {
				ID        uint64    `gorm:"primary_key;not null;comment:id"`
				Name      string    `gorm:"type:varchar(255);not null;default:'';comment:name"`
				Email     string    `gorm:"type:varchar(255);not null;default:'';uniqueIndex;comment:email"`
				Avatar    string    `gorm:"type:varchar(255);not null;default:'';comment:avatar"`
				Gender    int       `gorm:"type:tinyint(1);not null;default:0;comment:gender,1-female,2-male"`
				Birthday  time.Time `gorm:"type:date;default null;comment:birthday"`
				Salt      string    `gorm:"type:varchar(40);not null;default:'';comment:salt"`
				Password  string    `gorm:"type:varchar(100);not null;default:'';comment:password"`
				CreatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);comment:create time"`
				UpdatedAt time.Time `gorm:"type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:update time"`
			}

			return db.Migrator().AutoMigrate(&User{})
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable("users")
		},
	})
}
