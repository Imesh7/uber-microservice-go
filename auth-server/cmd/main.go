package main

import (
	"fmt"
	"uber-microservice-auth/api"
	"uber-microservice-auth/db"
	"uber-microservice-auth/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	connDB, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	v1 := router.Group("/v1")
	{
		v1.POST("/create-customer", handler.CreateCustomer)
		v1.POST("/create-driver", handler.CreateDriver)
		v1.POST("/customer-login", handler.LoginCustomer)
		v1.POST("/driver-login", handler.LoginDriver)
		v1.POST("/validate", handler.ValidateToken)
	}

	connDB.DB.AutoMigrate(&api.Customer{})
	connDB.DB.AutoMigrate(&api.Driver{})
	fmt.Println("application is startings..............")
	router.Run(":8001")
}
