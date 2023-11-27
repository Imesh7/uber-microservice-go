package handler

import (
	"net/http"
	"uber-microservice-trip-management/api"
	"github.com/gin-gonic/gin"
)

func CreateTripHandler(c *gin.Context) {
	var trip api.Trip
	err := c.Bind(&trip)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid data",
		})
		return
	}
	trip.CreateTrip(c)
}

//connect to customer websocket
func ConnectCustomerWS(c *gin.Context)  {
	api.ConnectCustomerToWebSocket(c)
}

//connect to driver websocket
func ConnectDriverWS(c *gin.Context)  {
	api.ConnectDriverToWebSocket(c)
}
