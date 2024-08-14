package db

func InitDB() {
	ConnectDB()
	InitRedis()
	InitQueryBuilder(GetWriteDB())
}
