package types


type OkResponseData struct {
	Status  int32 `json:"status"`
	Message string  `json:"message"`
	Data    map[string]string  `json:"data"`
}

type ErrorResponseData struct {
	Status       int32 `json:"status"`
	ErrorMessage string `json:"message"`
}
