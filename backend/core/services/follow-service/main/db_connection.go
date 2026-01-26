package main

import (
	"log/slog"
	"services/follow-service/internal"
	"services/pkg/common"
	"time"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db_init() *internal.FollowHandler {
	// Configure Destination
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:1234@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local"
	}
	slog.Info("DB Destination is : " + dsn)

	// DB Connection (MySQL)
	var db *gorm.DB
	var err error

	slog.Info("Start Connection...")
	for i := 0; i < 10; i++ {
		slog.Info("Try Connection... #", i+1)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err == nil {
			break
		}

		slog.Warn("DB Connection failed, Restart... ", "error", err)
		time.Sleep(time.Second)
	}

	if err != nil || db == nil {
		slog.Error("Eventualy DB Connection failed")
		os.Exit(1)
	}
	slog.Info("DB Connected! Start migration")

	if err := db.AutoMigrate(&internal.Follow{}); err != nil {
		slog.Error("Migration Failed!", "error", err)
		os.Exit(1)
	}

	slog.Info("Injecting Dependency")
	repo := &internal.FollowRepository{GormRepository: common.GormRepository[internal.Follow]{DB: db}}
	slog.Info("Repository is ready : " + repo.DB.Name())

	svc := &internal.FollowService{Repo: repo}
	slog.Info("Service is ready : " + svc.Repo.DB.Name())

	hdl := &internal.FollowHandler{Svc: svc}
	slog.Info("Handler is ready : " + hdl.Svc.Repo.DB.Name())

	return hdl
}
