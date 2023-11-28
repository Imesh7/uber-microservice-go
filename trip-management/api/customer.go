package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketConnWithId struct {
	TripId            string
	CustomerWebsocket *websocket.Conn
}

var CustomerList = make(map[string]*WebSocketConnWithId)

var customerUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ConnectCustomerToWebSocket(c *gin.Context) {
	websocketConn, err := customerUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer websocketConn.Close()
	query := c.Request.URL.Query()
	tripId := query.Get("tripId")
	customerSocket := &WebSocketConnWithId{
		TripId:            tripId,
		CustomerWebsocket: websocketConn,
	}
	CustomerList[tripId] = customerSocket

	//customer location if need update the db
	for {
		_, msg, error := websocketConn.ReadMessage()
		if error != nil {
			log.Fatal(error)
			return
		}
		fmt.Println("message received..............")
		fmt.Println(msg)
	}
}

func CustomerKafkaConsumer() {
	config := sarama.NewConfig()
	config.ClientID = "go-kafka-consumer"
	//host := os.Getenv("KAFKA_HOST")
	fmt.Fprintln(os.Stdout, []any{"jost is...........%s", "host"}...)
	kafkaHost := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))
	fmt.Println(kafkaHost)
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		fmt.Fprintln(os.Stdout, []any{"Errors is-- %s", err}...)
		log.Fatal(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Println("kafka consumer close error", err)
			//log.Fatal(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("driver-location", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
			//receive driver location from kafka
		case msg := <-partitionConsumer.Messages():

			valueString := string(msg.Value)
			fmt.Fprintln(os.Stdout, []any{"Received kafka driver location: %s", valueString}...)
			sendMessageToCustomer(msg.Value)
		case <-signals:
			return
		}
	}
}

func sendMessageToCustomer(msg []byte) {
	//to get the customer location using the APP or WEB
	fmt.Println("sending message is ", msg)
	var data map[string]string
	err := json.Unmarshal(msg, &data)
	if err != nil {
		fmt.Println("error in json unmarshaling", data["tripId"])
		return
	}
	fmt.Println("sending message is ", data["tripId"])
	//sending driver location details to customer client
	if CustomerList[data["tripId"]] != nil {
		err := CustomerList[data["tripId"]].CustomerWebsocket.WriteMessage(1, []byte(data["message"]))
		if err != nil {
			return
		}
	}
}
