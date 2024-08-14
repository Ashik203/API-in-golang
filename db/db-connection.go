package db

import (
	"app/config"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func connect(dbConfig config.DBConfig) *sqlx.DB {
	ConnStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	if !dbConfig.EnableSSLMode {
		ConnStr += " sslmode=disable"
	}

	dbCon, err := sqlx.Connect("postgres", ConnStr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dbCon.SetConnMaxIdleTime(
		time.Duration(dbConfig.MaxIdleTimeInMinute * int(time.Minute)),
	)

	return dbCon
}

func ConnectDB() {
	conf := config.GetConfig()
	
	readDb = connect(conf.DB.Read)
	if readDb != nil {
		slog.Info("Connected to read database!")
	}

	WriteDb = connect(conf.DB.Write)
	if WriteDb != nil {
		slog.Info("Connected to write database!")
	}
}

func CloseDB() {
	if err := readDb.Close(); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Disconnected from read database!")

	if err := WriteDb.Close(); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Disconnected from write database!")
}
