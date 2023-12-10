package main

import (
	"fmt"
	"log"
	"net"
	"uber-microservice-auth/api"
	"uber-microservice-auth/db"
	customerPb "uber-microservice-auth/gen/customer/proto"

	driverPb "uber-microservice-auth/gen/driver/proto"

	//"uber-microservice-auth/handler"

	//"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// router := gin.Default()

	connDB, err := db.ConnectDatabase()
	connDB.DB.AutoMigrate(&api.Customer{})
	connDB.DB.AutoMigrate(&api.Driver{})
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()
	server := grpc.NewServer()
	customerPb.RegisterCustomerAuthServiceServer(server, &api.CustomerAuthServiceServer{})
	driverPb.RegisterDriverAuthServiceServer(server, &api.DriverAuthServiceServer{})
	fmt.Println("auth server-application is starting..............")
	server.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}

	// v1 := router.Group("/v1")
	// {
	// 	v1.POST("/create-customer", handler.CreateCustomer)
	// 	v1.POST("/create-driver", handler.CreateDriver)
	// 	v1.POST("/customer-login", handler.LoginCustomer)
	// 	v1.POST("/driver-login", handler.LoginDriver)
	// 	v1.POST("/validate", handler.ValidateToken)
	// }

	// fmt.Println("application is starting..............")
	// router.Run(":8001")
}
