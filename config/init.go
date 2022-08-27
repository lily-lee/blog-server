package config

func Init() {
	InitConfig()
	InitDB()
	InitRedis()
	InitSnowflake()
}
