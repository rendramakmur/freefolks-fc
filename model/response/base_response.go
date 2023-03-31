package model

type BaseResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

func CreateResponse(status int, data interface{}, err error) *BaseResponse {
	if err != nil {
		return &BaseResponse{Status: status, Data: data, Error: err.Error()}
	}

	return &BaseResponse{Status: status, Data: data}
}
