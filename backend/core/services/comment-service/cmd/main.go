package cmd

import (
	"services/comment-service/cmd"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB Connection, return handler
	hdl := cmd.db_init()

	// Gin init
	r := gin.Default()

	// v2 Group
	v2 := r.Group("/api/v2/comments")
	{
		v2.GET("/:id", hdl.GetCommentByID)
		v2.GET("", hdl.GetComments)
		v2.POST("/", hdl.CreateComment)
		v2.PUT("/:id", hdl.UpdateComment)
		v2.DELETE("/:id", hdl.DeleteComment)
	}

	// Server Init
	r.Run(":8080")
}
