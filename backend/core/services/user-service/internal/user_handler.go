package internal

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Svc *UserService
}

func (hdl *UserHandler) CreateUser(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := hdl.Svc.CreateUser(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, res)
}

func (hdl *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := hdl.Svc.GetUserByID(userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

func (hdl *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Query("username")

	user, err := hdl.Svc.GetUserByUsername(username)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

func (hdl *UserHandler) UpdateUser(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	res, err := hdl.Svc.UpdateUser(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, res)
}

func (hdl *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	isDone, err := hdl.Svc.DeleteUser(userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"isDone": isDone})
}
