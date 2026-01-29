package main

import (
	"log/slog"
	"os"
	"services/pkg/common"
	"services/post-service/internal"
)

func dbMigrate() *internal.PostHandler {
	db := common.DBConfig("post")
	// migration
	if err := db.AutoMigrate(&internal.Post{}); err != nil {
		slog.Error("Post Migration Failed!", "error", err)
		os.Exit(1)
	}
	if err := db.AutoMigrate(&internal.PostImage{}); err != nil {
		slog.Error("PostImage Migration Failed!", "error", err)
		os.Exit(1)
	}
	if err := db.AutoMigrate(&internal.Like{}); err != nil {
		slog.Error("Like Migration Failed!", "error", err)
		os.Exit(1)
	}

	// Dependency Injection
	slog.Info("Injecting Dependency")
	repo := &internal.PostRepository{
		PostRepo:  common.GormRepository[internal.Post]{DB: db},
		ImageRepo: common.GormRepository[internal.PostImage]{DB: db},
		LikeRepo:  common.GormRepository[internal.Like]{DB: db},
	}
	slog.Info("Repository is ready : ", "db_name", repo.PostRepo.DB.Name())

	svc := &internal.PostService{Repo: repo}
	slog.Info("Service is ready : ", "db_name", svc.Repo.PostRepo.DB.Name())

	hdl := &internal.PostHandler{Svc: svc}
	slog.Info("Handler is ready : ", "db_name", hdl.Svc.Repo.PostRepo.DB.Name())

	return hdl
}
