package common

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Error   error       `json:"error"`
}
