package handler

import (
	"net/http"
	"uber-microservice-auth/api"
	"uber-microservice-auth/config"

	"github.com/gin-gonic/gin"
)

// create customer
func CreateCustomer(c *gin.Context) {
	var createUser config.CreateCustomerRequest

	err := c.ShouldBind(&createUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please fill the required fields",
		})
		return
	} else {
		//api.Create(c, &createUser)
	}
}

// create driver
func CreateDriver(c *gin.Context) {
	var createDriver config.CreateDriverRequest

	err := c.ShouldBind(&createDriver)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please fill the required fields",
		})
		return
	} else {
		//api.CreateDriver(c, &createDriver)
	}
}

// login customer
func LoginCustomer(c *gin.Context) {
	var loginCustomer config.LoginUserRequest

	err := c.ShouldBind(&loginCustomer)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please fill the required fields",
		})
	} else {
		//api.LoginUser(c, &loginCustomer)
	}
}

// login driver
func LoginDriver(c *gin.Context) {
	var loginDriver config.LoginUserRequest

	err := c.ShouldBind(&loginDriver)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Please fill the required fields",
		})
	} else {
		//api.LoginDriver(c, &loginDriver)
	}
}

func ValidateToken(c *gin.Context) {
	api.ValidateTokenClaims(c)
}
