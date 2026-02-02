package common

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		slog.Info("HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", time.Since(time.Now()),
			"ip", c.ClientIP(),
		)

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			slog.Error("Error Occured", "error", err)

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(404, gin.H{"message": "Not Found"})
			} else {
				c.JSON(500, gin.H{"message": "Internal Server Error"})
			}
		}
	}
}

func CommitError(fileName, functionName, message string, err error) error {
	// return nil, fmt.Errorf("CommentService.GetCommentsByID, "+
	// 		"error occured during converting string to uint: %w", err)
	slog.Error(message)
	return fmt.Errorf("%s.%s, %s: %w", fileName, functionName, message, err)
}
