package internal

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Svc *UserService
}

/*
* User Handler
 */

// CreateUser godoc
// @Summary      Sign In
// @Description  Create a new User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      UserRequest  true  "User DTO"
// @Success      201   {object}  models.UserResponse
// @Failure      400   {object}  map[string]string "Wrong Request"
// @Router       /users [post]
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

// GetUserByID godoc
// @Summary      Read Single User
// @Description  Read specific user with User ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  models.UserResponse
// @Router       /users/{id} [get]
func (hdl *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := hdl.Svc.GetUserByID(userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

// GetUserByUsername godoc
// @Summary      Read Single User with Username
// @Description  Read User by username
// @Tags         users
// @Produce      json
// @Param        username  query     string  true  "username"
// @Success      200       {object}  models.UserResponse
// @Router       /users [get]
func (hdl *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Query("username")

	user, err := hdl.Svc.GetUserByUsername(username)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, user)
}

// UpdateUser godoc
// @Summary      Update User
// @Description  Update the User
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      UserRequest  true  "User DTO"
// @Success      200   {object}  models.UserResponse
// @Router       /users [put]
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

// DeleteUser godoc
// @Summary      Delete User
// @Description  Delete User by User ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "id"
// @Success      200  {object}  map[string]bool "isDone: true"
// @Router       /users/{id} [delete]
func (hdl *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	isDone, err := hdl.Svc.DeleteUser(userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{"isDone": isDone})
}
