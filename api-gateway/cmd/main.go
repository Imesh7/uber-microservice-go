package main

import (
	"fmt"
	"log"
	"net/http"

	customer "api-gateway/gen/customer"
	driver "api-gateway/gen/driver"
	"api-gateway/handler"
	customer_service "api-gateway/service/customer"
	driver_service "api-gateway/service/driver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("application started.....")
	conn, err := grpc.Dial("auth-server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("got a GRPC error.....................")
		log.Println(err)
	}
	defer conn.Close()

	//grpc clients
	customerClient := customer.NewCustomerAuthServiceClient(conn)
	driverClient := driver.NewDriverAuthServiceClient(conn)

	//injects grpc clients to services
	customerService := customer_service.NewCustomerServiceImpl(customerClient)
	driverService := driver_service.NewDriverServiceImpl(driverClient)

	http.HandleFunc("/customer-login", func(w http.ResponseWriter, r *http.Request) {
		handler.CustomerLogin(w, r, customerService)
	})
	http.HandleFunc("/create-customer", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateCustomer(w, r, customerService)
	})

	http.HandleFunc("/driver-login", func(w http.ResponseWriter, r *http.Request) {
		handler.DriverLogin(w, r, driverService)
	})
	http.HandleFunc("/create-driver", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateDriver(w, r, driverService)
	})

	log.Println("Starting API Gateway server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
