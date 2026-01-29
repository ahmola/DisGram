package main

import (
	"log/slog"
	"services/follow-service/internal"
	"services/pkg/common"

	"os"
)

func dbInit() *internal.FollowHandler {
	// Configure Destination
	db := common.DBConfig("follow")

	if err := db.AutoMigrate(&internal.Follow{}); err != nil {
		slog.Error("Migration Failed!", "error", err)
		os.Exit(1)
	}

	slog.Info("Injecting Dependency")
	repo := &internal.FollowRepository{GormRepository: common.GormRepository[internal.Follow]{DB: db}}
	slog.Info("Repository is ready : ", "db_name", repo.DB.Name())

	svc := &internal.FollowService{Repo: repo}
	slog.Info("Service is ready : ", "db_name", svc.Repo.DB.Name())

	hdl := &internal.FollowHandler{Svc: svc}
	slog.Info("Handler is ready : ", "db_name", hdl.Svc.Repo.DB.Name())

	return hdl
}
