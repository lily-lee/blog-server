package config

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var Conf conf

type conf struct {
	Env           string `env:"ENV" envDefault:"develop"`
	Port          int    `env:"PORT" envDefault:"3000"`
	RedisAddr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	RedisPoolSize int    `env:"REDIS_POOL_SIZE" envDefault:"10"`
	DBName        string `env:"DB_NAME" envDefault:"blog"`
	DBHost        string `env:"DB_HOST" envDefault:"localhost"`
	DBPort        int    `env:"DB_PORT" envDefault:"3306"`
	DBUsername    string `env:"DB_USERNAME"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBIdlConn     int    `env:"DB_IDL_CONN" envDefault:"10"`
	DBMaxConn     int    `env:"DB_MAX_CONN" envDefault:"100"`

	SignKey       string        `env:"SIGN_KEY"`
	TokenDuration time.Duration `env:"TOKEN_DURATION" envDefault:"24h"`
}

func InitConfig(files ...string) {
	if len(os.Getenv("ENV")) == 0 {
		if err := godotenv.Load(files...); err != nil {
			panic(err)
		}
	}
	if err := env.Parse(&Conf); err != nil {
		panic(err)
	}

	fmt.Println("parse env success.")
}
