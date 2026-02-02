package common

import (
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConfig(serviceName string) *gorm.DB {
	// Configure Destination
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:1234@tcp(127.0.0.1:3306)/" + serviceName + "?useUnicode=true&characterEncoding=utf8mb4&serverTimezone=UTC"
	}
	slog.Info("DB Destination is : " + dsn)

	// DB connection (MySQL)

	var DB *gorm.DB
	var err error

	slog.Info("Start Connecting...")
	for i := 0; i < 10; i++ {
		slog.Info("Try Connection... #", i+1)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			break
		}

		slog.Warn("DB Connection failed, Restart... ", "error", err)
		time.Sleep(time.Second)
	}

	if err != nil || DB == nil {
		slog.Error("Eventually DB Connection failed.")
		os.Exit(1)
	}
	slog.Info(serviceName + " DB Connected! Start migrate")

	return DB
}
