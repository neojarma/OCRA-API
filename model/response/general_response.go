package response

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type EmptyObjectDataResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}

type EmptyArrayDataResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    struct{} `json:"data"`
}
