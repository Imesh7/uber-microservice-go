package handler

import (
	customer_service "api-gateway/service/customer"
	"net/http"
)

func CustomerLogin(w http.ResponseWriter, r *http.Request, customerService customer_service.CustomerService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	customerService.Login(w, r)
	return
}

func CreateCustomer(w http.ResponseWriter, r *http.Request, customerService customer_service.CustomerService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	customerService.CreateCustomer(w, r)
	return
}
