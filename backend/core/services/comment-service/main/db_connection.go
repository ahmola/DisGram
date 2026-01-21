package cmd

import (
	"services/comment-service/internal"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db_init() *internal.CommentHandler {
	// db connection (MySQL)
	dsn := os.Getenv("DB_DSN")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&internal.Comment{})

	// Dependency Injection
	repo := &internal.CommentRepository{DB: db}
	svc := &internal.CommentService{Repo: repo}
	hdl := &internal.CommentHandler{Service: svc}

	return hdl
}
