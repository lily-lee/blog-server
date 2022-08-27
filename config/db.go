package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := DSN{
		Username: Conf.DBUsername,
		Password: Conf.DBPassword,
		Host:     Conf.DBHost,
		Port:     Conf.DBPort,
		DbName:   Conf.DBName,
	}
	var err error
	DB, err = gorm.Open(mysql.Open(dsn.String()))
	if err != nil {
		panic(err)
	}

	pool, err := DB.DB()
	if err != nil {
		panic(err)
	}
	pool.SetConnMaxLifetime(time.Hour)
	pool.SetMaxIdleConns(Conf.DBIdlConn)
	pool.SetMaxOpenConns(Conf.DBMaxConn)

	if Conf.Env == "develop" {
		DB = DB.Debug()
	}
}

type DSN struct {
	Username  string
	Password  string
	Host      string
	Port      int
	DbName    string
	Charset   string
	ParseTime string
	Locale    string
}

func (d DSN) String() string {
	if d.Locale == "" {
		d.Locale = "Local"
	}

	if d.Charset == "" {
		d.Charset = "utf8mb4"
	}

	if d.ParseTime == "" {
		d.ParseTime = "True"
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DbName,
		d.Charset,
		d.ParseTime,
		d.Locale,
	)
}
