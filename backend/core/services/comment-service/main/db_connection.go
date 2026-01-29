package main

import (
	"log/slog"
	"services/comment-service/internal"
	"services/pkg/common"

	"os"
)

func dbInit() *internal.CommentHandler {
	// Configure Destination
	db := common.DBConfig("comment")

	if err := db.AutoMigrate(&internal.Comment{}); err != nil {
		slog.Error("Migration Failed!", "error", err)
		os.Exit(1)
	}

	// Dependency Injection
	slog.Info("Injecting Dependency")
	repo := &internal.CommentRepository{GormRepository: common.GormRepository[internal.Comment]{DB: db}}
	slog.Info("Repository is ready : ", "db_name", repo.DB.Name())

	svc := &internal.CommentService{Repo: repo}
	slog.Info("Service is ready : ", "db_name", svc.Repo.DB.Name())

	hdl := &internal.CommentHandler{Service: svc}
	slog.Info("Handler is ready : ", "db_name", hdl.Service.Repo.DB.Name())

	return hdl
}
