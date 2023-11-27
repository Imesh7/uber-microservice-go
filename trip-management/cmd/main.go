package main

import (
	"fmt"
	"uber-microservice-trip-management/api"
	"uber-microservice-trip-management/database"
	"uber-microservice-trip-management/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	go api.CustomerKafkaConsumer()
	go api.KafkaProducerConnect()
	router.POST("create-trip", handler.CreateTripHandler)
	router.GET("customer-websocket", handler.ConnectCustomerWS)
	router.GET("driver-websocket", handler.ConnectDriverWS)
	fmt.Println("Trip management application statrted............")
	mongoClient := database.ConnectDatabase()
	mongoDatabase := database.AccessDatabase(mongoClient, "uber-trip-db")
	database.AccessTripsCollection(mongoDatabase)
	router.Run(":8002")
}
