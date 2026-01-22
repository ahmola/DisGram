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
		c.Error(err)
		return
	}

	res, err := hdl.Service.CreateComment(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, res)
}

func (hdl *CommentHandler) GetComments(c *gin.Context) {
	postID := c.Query("postID")

	comments, err := hdl.Service.GetCommentsByPostID(postID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, comments)
}

func (hdl *CommentHandler) GetCommentByID(c *gin.Context) {
	commentID := c.Param("id")

	comment, err := hdl.Service.GetCommentByID(commentID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, comment)
}

func (hdl *CommentHandler) UpdateComment(c *gin.Context) {
	var req CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := hdl.Service.UpdateComment(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}

func (hdl *CommentHandler) DeleteComment(c *gin.Context) {
	commentID := c.Param("id")

	isDone, err := hdl.Service.DeleteComment(commentID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"isDone": isDone})
}
