package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"uber-microservice-trip-management/database"

	"github.com/gin-gonic/gin"
)

//import "time"z

type Trip struct {
	TripId int32 `bson:"trip_ids" json:"trip_Id"`
	// CustomerId                int32
	// PickedUpLocationLongitude float32
	// PickedUpLocationLatitude  float32
	// TripCost                  int32
	// EndLocationLongitude      float32
	// EndLocationLatitude       float32
	// TripStatus                TripStatus
	// CreatedAt                 time.Time
}

type TripStatus int16

const (
	PendingTrip TripStatus = iota + 1
	TripAccepted
	TripPicked
	TripCompleted
	TripCancelled
)

func (trip *Trip) CreateTrip(c *gin.Context) {
	result, error := database.MongoCollection.TripsCollection.InsertOne(context.TODO(), &trip)
	if error != nil {
		log.Println(error)
	}
	fmt.Print(result.InsertedID)
	c.JSON(http.StatusOK, gin.H{
		"message": "success ",
	})
}

func pushTripToDrivers() {

}

func stopPushToDrivers() {

}

func CancelTrip() {

}
