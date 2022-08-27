package config

import (
	"time"

	"github.com/sony/sonyflake"
)

var SnowFlake *sonyflake.Sonyflake

func InitSnowflake() {
	SnowFlake = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: time.Now(),
	})
}
