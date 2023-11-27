package main

import (
	"fmt"
	"uber-microservice-auth/api"
	"uber-microservice-auth/db"
	"uber-microservice-auth/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("application is startings..............")
	router := gin.Default()

	connDB, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	v1 := router.Group("/v1")
	{
		v1.POST("/create-user", handler.CreateUser)
		v1.POST("/login", handler.LoginUser)
		v1.POST("/validate", handler.ValidateToken)
	}

	connDB.DB.AutoMigrate(&api.User{})
	router.Run(":8001")
}
