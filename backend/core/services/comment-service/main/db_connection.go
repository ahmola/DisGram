package main

import (
	"log/slog"
	"services/comment-service/internal"
	"services/pkg/common"
	"time"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db_init() *internal.CommentHandler {
	// Configure Destination
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:1234@tcp(127.0.0.1:3306)/comment?charset=utf8mb4&parseTime=True&loc=Local"
	}
	slog.Info("DB Destination is : " + dsn)

	// db connection (MySQL)

	var db *gorm.DB
	var err error

	slog.Info("Start Connecting...")
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
		slog.Error("Eventually DB Connection failed.")
		os.Exit(1)
	}
	slog.Info("DB Connected! Start migrate")

	if err := db.AutoMigrate(&internal.Comment{}); err != nil {
		slog.Error("Migration Failed!", "error", err)
		os.Exit(1)
	}

	// Dependency Injection
	slog.Info("Injecting Dependency")
	repo := &internal.CommentRepository{GormRepository: common.GormRepository[internal.Comment]{DB: db}}
	slog.Info("Repository is ready : " + repo.DB.Name())

	svc := &internal.CommentService{Repo: repo}
	slog.Info("Service is ready : " + svc.Repo.DB.Name())

	hdl := &internal.CommentHandler{Service: svc}
	slog.Info("Handler is ready : " + hdl.Service.Repo.DB.Name())

	return hdl
}
