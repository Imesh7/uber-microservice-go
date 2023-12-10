package customer_service

import "net/http"

type CustomerService interface {
	Login(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
}