package handler

import (
	"net/http"
	"uber-microservice-auth/api"
	"uber-microservice-auth/config"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var createUser config.CreateUserRequest

	err := c.ShouldBind(&createUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please fill the required fields",
		})
		return
	} else {
		api.CreatUser(c, &createUser)
	}
}

func LoginUser(c *gin.Context) {
	var loginUser config.LoginUserRequest

	err := c.ShouldBind(&loginUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please fill the required fields",
		})
	} else {
		println(loginUser.Email)
		api.LoginUser(c, &loginUser)
	}
}

func ValidateToken(c *gin.Context) {
	api.ValidateTokenClaims(c)
}
