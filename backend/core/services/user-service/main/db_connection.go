package main

import (
	"log/slog"
	"services/pkg/common"
	"services/user-service/internal"

	"os"
)

func dbInit() *internal.UserHandler {
	// Configure Destination
	db := common.DBConfig("user")

	if err := db.AutoMigrate(&internal.User{}); err != nil {
		slog.Error("Migration Failed!", "error", err)
		os.Exit(1)
	}

	// Dependency Injection
	slog.Info("Injecting Dependency")
	repo := &internal.UserRepository{GormRepository: common.GormRepository[internal.User]{DB: db}}
	slog.Info("Repository is ready : ", "db_name", repo.DB.Name())

	svc := &internal.UserService{Repo: repo}
	slog.Info("Service is ready : ", "db_name", repo.DB.Name())

	hdl := &internal.UserHandler{Svc: svc}
	slog.Info("Handler is ready : ", "db_name", repo.DB.Name())

	return hdl
}
