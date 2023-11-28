package api

import (
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var driverProducer sarama.SyncProducer

var driverUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ConnectDriverToWebSocket(c *gin.Context) {
	websocketConn, err := driverUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(1000, gin.H{
			
			"message": "Cannot connect to the web socket",
		})
		return
	}
	for {
		//receive location update from the driver client
		_, msg, err := websocketConn.ReadMessage()
		if err != nil {
			c.JSON(1000, gin.H{
				"message": "Failed to read messaage",
			})
			return
		}
		fmt.Println(msg)
		//send the driver location data to customer client
		sendProducerMessage(msg)
	}

}

// send driver location to the customer via kafka
func KafkaProducerConnect() {
	con := sarama.NewConfig()
	con.Producer.Return.Successes = true
	//kafkaHost := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, con)
	if err != nil {
		fmt.Println("Cannot connect to producer")
		log.Fatal(err)
	}
	driverProducer = producer

	// defer func() {
	// 	if err := producer.Close(); err != nil {
	// 		fmt.Println("kafka producter close error", err)
	// 		//log.Fatal(err)
	// 	}
	// }()
}

func sendProducerMessage(msg []byte) {
	message := &sarama.ProducerMessage{
		Topic: "driver-location",
		Value: sarama.StringEncoder(msg),
	}
	partition, offset, err := driverProducer.SendMessage(message)
	if err != nil {
		fmt.Fprintln(os.Stdout, []any{"message Error................. %s", err}...)
		log.Fatal(err)
	}
	log.Printf("Produced message to partition %d at offset %d\n", partition, offset)
}

func ReceiveDriverLocationFromClient() {

}

func AcceptTrip() {

}

func FinishTrip() {

}

func DeclineTrip() {

}
