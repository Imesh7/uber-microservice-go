package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"uber-microservice-trip-management/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Trip struct {
	TripId                    primitive.ObjectID `bson:"trip_id,omitempty" json:"trip_id"`
	CustomerId                int32              `bson:"customer_id" json:"customer_id"`
	PickedUpLocationLongitude float32            `bson:"picked_up_location_longitude" json:"picked_up_location_longitude"`
	PickedUpLocationLatitude  float32            `bson:"picked_up_location_latitude" json:"picked_up_location_latitude"`
	TripCost                  float32            `bson:"trip_cost" json:"trip_cost"`
	EndLocationLongitude      float32            `bson:"end_location_longitude" json:"end_location_longitude"`
	EndLocationLatitude       float32            `bson:"end_location_latitude" json:"end_location_latitude"`
	tripStatus                TripStatus         `bson:"trip_status"`
	createdAt                 time.Time          `bson:"created_at"`
}

type TripStatus int16

const (
	TripCreated TripStatus = iota + 1
	SearchingDriver
	TripAccepted
	TripPicked
	TripCompleted
	TripCancelled
)

func (trip *Trip) CreateTrip(c *gin.Context) {
	trip.createdAt = time.Now()
	trip.tripStatus = TripCreated

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
