package db

import "app/redis"

func InitDB() {
	ConnectDB()
	redis.InitRedis()
	InitQueryBuilder()
	InitUsersRepo()
	InitBookRepo()
}
