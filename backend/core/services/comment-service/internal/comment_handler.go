package internal

import (
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	Service *CommentService
}

// CreateComment godoc
// @Summary      Create Comment
// @Description  Create new comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        comment  body      CommentRequest  true  "Comment DTO"
// @Success      201      {object}  CommentResponse
// @Failure      400      {object}  map[string]string "잘못된 요청"
// @Router       /comments [post]
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

// GetComments godoc
// @Summary      Read the list of comments by post
// @Description  Gets all comments on a particular post by postID.
// @Tags         comments
// @Produce      json
// @Param        postID   query     string  true  "post ID"
// @Success      200      {array}   Comment
// @Router       /comments [get]
func (hdl *CommentHandler) GetComments(c *gin.Context) {
	postID := c.Query("postID")

	comments, err := hdl.Service.GetCommentsByPostID(postID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, comments)
}

// GetCommentByID godoc
// @Summary      Read Single Comment
// @Description  Use the comment ID to query the details of a particular comment
// @Tags         comments
// @Produce      json
// @Param        id   path      string  true  "댓글 ID"
// @Success      200  {object}  Comment
// @Router       /comments/{id} [get]
func (hdl *CommentHandler) GetCommentByID(c *gin.Context) {
	commentID := c.Param("id")

	comment, err := hdl.Service.GetCommentByID(commentID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, comment)
}

// UpdateComment godoc
// @Summary      Modify Comment
// @Description  Modify the content of an existing comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        comment  body      CommentRequest  true  "Comment DTO"
// @Success      200      {object}  Comment
// @Router       /comments [put]
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

// DeleteComment godoc
// @Summary      Delete Comment
// @Description  Use the comment ID to delete a specific comment
// @Tags         comments
// @Produce      json
// @Param        id   path      string  true  "Comment ID"
// @Success      200  {object}  map[string]bool "isDone: true"
// @Router       /comments/{id} [delete]
func (hdl *CommentHandler) DeleteComment(c *gin.Context) {
	commentID := c.Param("id")

	isDone, err := hdl.Service.DeleteComment(commentID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"isDone": isDone})
}
