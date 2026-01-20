package cmd

import (
	"services/comment-service/internal"

	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db_init(){
	// db connection (MySQL)
	dsn := os.Getenv("DB_DSN")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&internal.Comment{})

	// Dependency Injection
	repo := &internal.CommentRepository{DB: db}
	svc := &service.CommentService{Repo: repo}
	hdl := &handler.CommentHandler{Service: svc}

	return hdl
}
