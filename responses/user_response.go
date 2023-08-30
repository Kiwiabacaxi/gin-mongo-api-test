package responses

type UserResponse struct {
	Status  int                    `json:"status"`  // 200, 400, 404, 500
	Message string                 `json:"message"` // "Success", "Bad Request", "Not Found", "Internal Server Error"
	Data    map[string]interface{} `json:"data"`    // map[string]interface{} tipo genérico
}

