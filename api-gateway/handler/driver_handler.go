package handler

import (
	"api-gateway/service/driver"
	"net/http"
)

func DriverLogin(w http.ResponseWriter, r *http.Request,driverService driver_service.DriverService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	driverService.Login(w,r)
	return
}

func CreateDriver(w http.ResponseWriter, r *http.Request,driverService driver_service.DriverService) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	driverService.CreateDriver(w,r)
	return
}