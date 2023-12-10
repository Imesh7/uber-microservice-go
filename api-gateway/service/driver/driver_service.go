package driver_service

import "net/http"

type DriverService interface {
	Login(w http.ResponseWriter, r *http.Request)
	CreateDriver(w http.ResponseWriter, r *http.Request)
}