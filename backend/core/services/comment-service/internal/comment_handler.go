package internal

import (
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	Service *CommentService
}

func (hdl *CommentHandler) CreateComment(c *gin.Context) {
	var req CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := hdl.Service.CreateComment(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, res)
}

func (hdl *CommentHandler) GetComments(c *gin.Context) {
	postID := c.Query("postID")

	comments, err := hdl.Service.GetCommentsByPostID(postID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, comments)
}

func (hdl *CommentHandler) GetCommentByID(c *gin.Context) {

}

func (hdl *CommentHandler) UpdateComment(c *gin.Context) {

}

func (hdl *CommentHandler) DeleteComment(c *gin.Context) {

}
