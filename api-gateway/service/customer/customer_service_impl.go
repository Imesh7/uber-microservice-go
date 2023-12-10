package customer_service

import (
	"api-gateway/gen/customer"
	"api-gateway/types"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type customerServiceImpl struct {
	client customer.CustomerAuthServiceClient
}

func NewCustomerServiceImpl(customerAuthClient customer.CustomerAuthServiceClient) CustomerService {
	return &customerServiceImpl{
		client: customerAuthClient,
	}
}

func (c *customerServiceImpl) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var loginRequestProto customer.LoginRequestProto
	if err := json.Unmarshal(body, &loginRequestProto); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	res, err := c.client.Login(context.Background(), &loginRequestProto)
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

func (c *customerServiceImpl) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var createRequestProto customer.CreateCustomerRequestProto
	if err := json.Unmarshal(body, &createRequestProto); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	res, err := c.client.Create(context.Background(), &createRequestProto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(res.StatusCode))
	json.NewEncoder(w).Encode(res)
	return
}
