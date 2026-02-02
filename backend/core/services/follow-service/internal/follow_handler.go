package internal

import (
	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	Svc *FollowService
}

// CreateFollow godoc
// @Summary      Create Follow
// @Description  Follow specific User
// @Tags         follows
// @Accept       json
// @Produce      json
// @Param        follow  body      FollowRequest  true  "Follow DTO"
// @Success      201     {object}  FollowResponse
// @Failure      400     {object}  map[string]string "Wrong Request"
// @Router       /follows [post]
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

// GetFolloweesByFollowerID godoc
// @Summary      List of Users followed by specific User
// @Description  Get All the Users' ID followed by specific User with User ID
// @Tags         follows
// @Produce      json
// @Param        follower_id  query     string  true  "Followees' ID"
// @Success      200          {array}   uint
// @Router       /follows/followees [get]
func (hdl *FollowHandler) GetFolloweesByFollowerID(c *gin.Context) {
	followerID := c.Query("follower_id")

	followees, err := hdl.Svc.GetFolloweesByFollowerID(followerID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, followees)
}

// GetFollowersByFolloweeID godoc
// @Summary      List of Users following specific User
// @Description  Get all the Users' ID following specific User
// @Tags         follows
// @Produce      json
// @Param        followee_id  query     string  true  "id of user who are followed"
// @Success      200          {array}   uint
// @Router       /follows/followers [get]
func (hdl *FollowHandler) GetFollowersByFolloweeID(c *gin.Context) {
	followeeID := c.Query("followee_id")

	followers, err := hdl.Svc.GetFollowersByFolloweeID(followeeID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, followers)
}

// DeleteFollow godoc
// @Summary      Cancel Follow
// @Description  Delete the Follow Relationship
// @Tags         follows
// @Produce      json
// @Param        id   path      string  true  "Follow ID"
// @Success      200  {object}  map[string]bool "isDone: true"
// @Router       /follows/{id} [delete]
func (hdl *FollowHandler) DeleteFollow(c *gin.Context) {
	followID := c.Param("id")

	isDone, err := hdl.Svc.DeleteFollow(followID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"isDone": isDone})
}
