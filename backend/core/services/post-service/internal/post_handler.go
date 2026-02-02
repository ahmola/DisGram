package internal

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Svc *PostService
}

/*
* post handler
 */

// CreatePost godoc
// @Summary      Create Post
// @Description  Create a new Post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      PostRequest  true  "PostDTO"
// @Success      201   {object}  Post
// @Failure      400   {object}  map[string]string "Wrong Request"
// @Router       /posts [post]
func (hdl *PostHandler) CreatePost(c *gin.Context) {
	var req PostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := hdl.Svc.CreatePost(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, res)
}

// GetPostById godoc
// @Summary      Read Post Info
// @Description  Read Post and Images of the post
// @Tags         posts
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Success      200  {object}  Post
// @Router       /posts/{id} [get]
func (hdl *PostHandler) GetPostById(c *gin.Context) {
	postID := c.Param("id")
	slog.Info("GetPostById called: ", "postId", postID)

	res, err := hdl.Svc.GetPostById(postID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}

// UpdatePost godoc
// @Summary      Update Post
// @Description  Update the post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      PostRequest  true  "Update Post"
// @Success      200   {object}  Post
// @Router       /posts [put]
func (hdl *PostHandler) UpdatePost(c *gin.Context) {
	var req PostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := hdl.Svc.UpdatePost(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}

// DeletePost godoc
// @Summary      Delete Post
// @Description  Delete the post, images of the post, likes of the post
// @Tags         posts
// @Produce      json
// @Param        id   path      string  true  "게시글 ID"
// @Success      200  {object}  map[string]bool "성공 여부"
// @Router       /posts/{id} [delete]
func (hdl *PostHandler) DeletePost(c *gin.Context) {
	postID := c.Param("id")

	res, err := hdl.Svc.DeletePostById(postID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}

/*
* like handler
 */

// CreateLike godoc
// @Summary      Create Like
// @Description  User add Like on the post. Cannot be duplicated
// @Tags         likes
// @Accept       json
// @Produce      json
// @Param        like  body      LikeRequest  true  "Like DTO"
// @Success      201   {object}  LikeResponse
// @Router       /likes [post]
func (hdl *PostHandler) CreateLike(c *gin.Context) {
	var req LikeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := hdl.Svc.CreateLike(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, res)
}

// GetAllLikesByPostID godoc
// @Summary      List of Likes
// @Description  Read all the likes of the specific post.
// @Tags         likes
// @Produce      json
// @Param        postId  path      string  true  "post ID"
// @Success      200     {array}   LikeResponse
// @Router       /posts/likes/{postId} [get]
func (hdl *PostHandler) GetAllLikesByPostID(c *gin.Context) {
	postID := c.Param("postId")

	res, err := hdl.Svc.GetAllLikesByPostID(postID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}

// DeleteLike godoc
// @Summary      Cancel Like
// @Description  Delete the Like of Post by Like ID
// @Tags         likes
// @Produce      json
// @Param        id   path      string  true  "like ID"
// @Success      200  {object}  map[string]bool "success"
// @Router       /likes/{id} [delete]
func (hdl *PostHandler) DeleteLike(c *gin.Context) {
	likeID := c.Param("id")

	res, err := hdl.Svc.DeleteLike(likeID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}
