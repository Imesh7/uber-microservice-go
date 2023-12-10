package driver_service

import (
	"api-gateway/gen/driver"
	"api-gateway/types"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type driverServiceImpl struct {
	conn driver.DriverAuthServiceClient
}

func NewDriverServiceImpl(client driver.DriverAuthServiceClient) DriverService {
	return &driverServiceImpl{
		conn: client,
	}
}

func (c *driverServiceImpl) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var loginRequestProto driver.LoginRequestProto
	if err := json.Unmarshal(body, &loginRequestProto); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	res, err := c.conn.Login(context.Background(), &loginRequestProto)
	if err != nil {
		errorResponse := types.ErrorResponseData{
			Status:       400,
			ErrorMessage: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.StatusCode))
	json.NewEncoder(w).Encode(res)
	return
}

func (c *driverServiceImpl) CreateDriver(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var createRequestProto driver.CreateDriverRequestProto
	if err := json.Unmarshal(body, &createRequestProto); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	res, err := c.conn.Create(context.Background(), &createRequestProto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.StatusCode))
	json.NewEncoder(w).Encode(res)
	return
}
