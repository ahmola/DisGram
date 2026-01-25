package internal

import (
	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	Svc *FollowService
}

func (hdl *FollowHandler) CreateFollow(c *gin.Context) {
	var req FollowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := hdl.Svc.CreateFollow(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, res)
}

func (hdl *FollowHandler) GetFolloweesByFollowerID(c *gin.Context) {
	followerID := c.Param("id")

	followees, err := hdl.Svc.GetFolloweesByFollowerID(followerID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, followees)
}

func (hdl *FollowHandler) GetFollowersByFolloweeID(c *gin.Context) {
	followeeID := c.Param("id")

	followers, err := hdl.Svc.GetFollowersByFolloweeID(followeeID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, followers)
}

func (hdl *FollowHandler) DeleteFollow(c *gin.Context) {
	followID := c.Param("id")

	isDone, err := hdl.Svc.DeleteFollow(followID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"isDone": isDone})
}
