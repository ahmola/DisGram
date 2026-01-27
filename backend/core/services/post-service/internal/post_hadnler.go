package internal

import "github.com/gin-gonic/gin"

type PostHandler struct {
	Svc *PostService
}

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

func (hdl *PostHandler) GetPostById(c *gin.Context) {
	postID := c.Param("id")

	res, err := hdl.Svc.GetPostById(postID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}

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

func (hdl *PostHandler) DeletePost(c *gin.Context) {
	postID := c.Param("id")

	res, err := hdl.Svc.DeletePostById(postID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}
